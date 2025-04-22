package repository

import (
	"context"
	localization "earnforglance/server/domain/localization"
	news "earnforglance/server/domain/news"
	domain "earnforglance/server/domain/public"
	security "earnforglance/server/domain/security"
	"earnforglance/server/service/data/mongo"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type newsItemRepository struct {
	database   mongo.Database
	collection string
}

func NewNewsItemRepository(db mongo.Database, collection string) domain.NewsItemRepository {
	return &newsItemRepository{
		database:   db,
		collection: collection,
	}
}

func (lr *newsItemRepository) GetNewsItems(c context.Context, filter domain.NewsItemRequest) ([]domain.NewsItemsResponse, error) {
	var result []domain.NewsItemsResponse
	var newsItems []news.NewsItem
	err := error(nil)

	idHex, err := primitive.ObjectIDFromHex(filter.ID)
	if err == nil {
		var newsItem news.NewsItem

		collection := lr.database.Collection(news.CollectionNewsItem)
		err = collection.FindOne(c, bson.M{"_id": idHex}).Decode(&newsItem)
		if err != nil {
			return result, err
		}

		item, err := PrepareNewsItem(lr, c, newsItem, filter)
		if err != nil {
			return result, err
		}

		result = append(result, domain.NewsItemsResponse{News: []domain.NewsItemResponse{item}})
		return result, err
	}

	query := bson.M{"published": true}

	if filter.AllowComments {
		query["allow_comments"] = filter.AllowComments
	}

	sortOrder := 1
	if filter.Sort == "desc" {
		sortOrder = -1
	}

	for _, value := range filter.Filters {

		if value.Operator == "contains" {
			query[value.Field] = bson.M{"$regex": value.Value, "$options": "i"}
		} else if value.Operator == "not_contains" {
			query[value.Field] = bson.M{"$not": bson.M{"$regex": value.Value, "$options": "i"}}
		} else if value.Operator == "mayor_than" {
			query[value.Field] = bson.M{"$gt": value.Value}
		} else if value.Operator == "minor_than" {
			query[value.Field] = bson.M{"$lt": value.Value}
		} else {
			query[value.Field] = value.Value
		}
	}

	findOptions := options.Find().
		SetSort(bson.D{{Key: "created_on_utc", Value: sortOrder}}).
		SetLimit(int64(filter.Limit))

	collection := lr.database.Collection(news.CollectionNewsItem)
	cursor, err := collection.Find(c, query, findOptions)
	if err != nil {
		return result, err
	}

	err = cursor.All(c, &newsItems)
	if err != nil {
		return result, err
	}

	var items []domain.NewsItemResponse
	for i := range newsItems {
		item, err := PrepareNewsItem(lr, c, newsItems[i], filter)
		if err != nil {
			return result, err
		}
		items = append(items, item)
	}

	result = append(result, domain.NewsItemsResponse{News: items})

	return result, err
}

func PrepareNewsItem(nr *newsItemRepository, c context.Context, newsItem news.NewsItem, filter domain.NewsItemRequest) (domain.NewsItemResponse, error) {
	var result domain.NewsItemResponse
	err := error(nil)

	for i := range filter.Content {
		switch filter.Content[i] {
		case "comments":
			result.Comments, err = PrepareNewsComments(nr, c, newsItem)
		}
	}

	if filter.Lang != "" {
		result.News, err = PrepareNewsItemLang(nr, c, newsItem, filter.Lang)
	} else {
		result.News = newsItem

	}

	return result, err
}

func PrepareNewsComments(nr *newsItemRepository, c context.Context, newsItem news.NewsItem) ([]news.NewsComment, error) {
	var result []news.NewsComment
	err := error(nil)
	query := bson.M{"is_approved": true, "news_item_id": newsItem.ID}

	collection := nr.database.Collection(news.CollectionNewsComment)
	cursor, err := collection.Find(c, query)

	if err != nil {
		return result, err
	}

	err = cursor.All(c, &result)
	if err != nil {
		return result, err
	}

	return result, err
}

func PrepareNewsItemLang(r *newsItemRepository, c context.Context, item news.NewsItem, lang string) (news.NewsItem, error) {
	var itemLang = item

	locale, err := GetLangugaNewsItemByCode(r, c, lang)
	if err != nil {
		return itemLang, err
	}

	record, err := GetRecordaNewsItemByCode(r, c, news.CollectionNewsItem)
	if err != nil {
		return itemLang, err
	}

	var items []localization.LocalizedProperty
	collection := r.database.Collection(localization.CollectionLocalizedProperty)
	cursor, err := collection.Find(c, bson.M{"entity_id": record.ID, "language_id": locale.ID, "locale_key_group": item.ID.Hex()})

	if err != nil {
		return itemLang, err
	}

	err = cursor.All(c, &items)
	if err != nil {
		return itemLang, err
	}

	for i := range items {
		switch items[i].LocaleKey {
		case "title":
			itemLang.Title = items[i].LocaleValue
		case "short":
			itemLang.Short = items[i].LocaleValue
		case "full":
			itemLang.Full = items[i].LocaleValue
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

func GetLangugaNewsItemByCode(nr *newsItemRepository, c context.Context, lang string) (localization.Language, error) {
	collection := nr.database.Collection(localization.CollectionLanguage)
	var item localization.Language
	err := collection.FindOne(c, bson.M{"unique_seo_code": lang}).Decode(&item)
	if err != nil {
		return item, err
	}
	return item, err
}

func GetRecordaNewsItemByCode(nr *newsItemRepository, c context.Context, name string) (security.PermissionRecord, error) {
	collection := nr.database.Collection(security.CollectionPermissionRecord)
	var item security.PermissionRecord
	err := collection.FindOne(c, bson.M{"system_name": name}).Decode(&item)
	if err != nil {
		return item, err
	}
	return item, err
}
