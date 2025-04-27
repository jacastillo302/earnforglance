package repository

import (
	"context"
	localization "earnforglance/server/domain/localization"
	domain "earnforglance/server/domain/public"
	topicdomain "earnforglance/server/domain/topics"
	"earnforglance/server/service/data/mongo"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type topicRepository struct {
	database   mongo.Database
	collection string
}

func NewTopicRepository(db mongo.Database, collection string) domain.TopicRepository {
	return &topicRepository{
		database:   db,
		collection: collection,
	}
}

func (r *topicRepository) GetTopicSecret(c context.Context, filter domain.TopicRequest) (domain.TopicsResponse, error) {
	var result domain.TopicsResponse

	idHex, err := bson.ObjectIDFromHex(filter.ID)
	if err == nil {
		var topicitem topicdomain.Topic

		collection := r.database.Collection(topicdomain.CollectionTopic)
		err = collection.FindOne(c, bson.M{"_id": idHex, "published": true}).Decode(&topicitem)
		if err != nil {
			return result, err
		}

		item, err := PrepareTopicSecret(r, c, topicitem, filter.Content, filter.Lang, filter.Password)
		if err != nil {
			return result, err
		}

		result = domain.TopicsResponse{Topics: []domain.TopicResponse{item}}
		return result, err
	}

	return result, err
}

func (r *topicRepository) GetTopics(c context.Context, filter domain.TopicRequest) ([]domain.TopicsResponse, error) {

	var result []domain.TopicsResponse
	var topics []topicdomain.Topic

	idHex, err := bson.ObjectIDFromHex(filter.ID)
	if err == nil {
		var topicitem topicdomain.Topic

		collection := r.database.Collection(topicdomain.CollectionTopic)
		err = collection.FindOne(c, bson.M{"_id": idHex, "published": true}).Decode(&topicitem)
		if err != nil {
			return result, err
		}

		item, err := PrepareTopic(r, c, topicitem, filter.Content, filter.Lang)
		if err != nil {
			return result, err
		}

		result = append(result, domain.TopicsResponse{Topics: []domain.TopicResponse{item}})
		return result, err
	}

	if filter.Limit == 0 {
		filter.Limit = 5
	}

	sortOrder := 1
	if filter.Sort == "desc" {
		sortOrder = -1
	}

	query := bson.M{"published": true}

	if filter.IncludeInTopMenu {
		query["include_in_top_menu"] = filter.IncludeInTopMenu
	}

	for _, value := range filter.Filters {
		if value.Operator == "contains" {
			query[value.Field] = bson.M{"$regex": value.Value, "$options": "i"}
		} else if value.Operator == "not_contains" {
			query[value.Field] = bson.M{"$not": bson.M{"$regex": value.Value, "$options": "i"}}
		} else {
			query[value.Field] = value.Value
		}
	}

	limit := int64(filter.Limit)

	findOptions := options.Find().
		SetSort(bson.D{{Key: "_id", Value: sortOrder}}).
		SetLimit(limit).
		SetProjection(bson.D{{Key: "password", Value: 0}})

	collection := r.database.Collection(topicdomain.CollectionTopic)
	cursor, err := collection.Find(c, query, findOptions)
	if err != nil {
		return result, err
	}

	err = cursor.All(c, &topics)
	if err != nil {
		return result, err
	}

	var items []domain.TopicResponse
	for i := range topics {
		item, err := PrepareTopic(r, c, topics[i], filter.Content, filter.Lang)
		if err != nil {
			return result, err
		}
		items = append(items, item)
	}

	result = append(result, domain.TopicsResponse{Topics: items})

	return result, err

}

func PrepareTopic(r *topicRepository, c context.Context, topic topicdomain.Topic, content []string, lang string) (domain.TopicResponse, error) {
	var result domain.TopicResponse
	err := error(nil)

	for i := range content {
		switch content[i] {
		case "template":
			result.Template, err = PrepareTopicTemplate(r, c, topic)
		}
	}

	if lang != "" {
		result.Topic, err = PrepareTopicLang(r, c, topic, lang)
	} else {
		result.Topic = topic

	}

	if result.Topic.IsPasswordProtected {
		result.Topic.Title = "This topic is password protected"
		result.Topic.Body = "This topic is password protected"
		result.Topic.Password = ""
	}

	return result, err
}

func PrepareTopicSecret(tr *topicRepository, c context.Context, topic topicdomain.Topic, content []string, lang string, password string) (domain.TopicResponse, error) {
	var result domain.TopicResponse

	err := error(nil)

	for i := range content {
		switch content[i] {
		case "template":
			result.Template, err = PrepareTopicTemplate(tr, c, topic)
		}
	}

	if lang != "" {
		result.Topic, err = PrepareTopicLang(tr, c, topic, lang)
	} else {
		result.Topic = topic
	}

	if result.Topic.IsPasswordProtected {
		if result.Topic.Password != password {
			result.Topic.Title = "this topic is password protected, the password is not correct"
			result.Topic.Body = "This topic is password protected, the password is not correct"
			result.Topic.Password = ""
		}
	}

	return result, err
}

func PrepareTopicLang(r *topicRepository, c context.Context, item topicdomain.Topic, lang string) (topicdomain.Topic, error) {
	var itemLang = item

	locale, err := GetLangugaByCode(c, lang, r.database.Collection(localization.CollectionLanguage))
	if err != nil {
		return itemLang, err
	}

	record, err := GetRercordBySystemName(c, topicdomain.CollectionTopic, r.database.Collection(topicdomain.CollectionTopicTemplate))
	if err != nil {
		return itemLang, err
	}

	items, err := GetLocalizedProperty(c, record.ID, locale.ID, item.ID.Hex(), r.database.Collection(localization.CollectionLocalizedProperty))
	if err != nil {
		return itemLang, err
	}

	for i := range items {
		switch items[i].LocaleKey {
		case "title":
			itemLang.Title = items[i].LocaleValue
		case "body":
			itemLang.Body = items[i].LocaleValue
		case "meta_title":
			itemLang.MetaTitle = items[i].LocaleValue
		case "meta_keywords":
			itemLang.MetaKeywords = items[i].LocaleValue
		case "meta_description":
			itemLang.MetaDescription = items[i].LocaleValue
		}
	}

	return itemLang, err
}

func PrepareTopicTemplate(tr *topicRepository, c context.Context, item topicdomain.Topic) (topicdomain.TopicTemplate, error) {
	var template topicdomain.TopicTemplate
	collection := tr.database.Collection(topicdomain.CollectionTopicTemplate)
	err := collection.FindOne(c, bson.M{"_id": item.TopicTemplateID}).Decode(&template)
	if err != nil {
		return template, err
	}
	return template, err
}
