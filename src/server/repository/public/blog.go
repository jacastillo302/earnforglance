package repository

import (
	"context"
	blog "earnforglance/server/domain/blogs"
	localization "earnforglance/server/domain/localization"
	domain "earnforglance/server/domain/public"
	security "earnforglance/server/domain/security"
	"earnforglance/server/service/data/mongo"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type blogRepository struct {
	database   mongo.Database
	collection string
}

func NewBlogRepository(db mongo.Database, collection string) domain.BlogRepository {
	return &blogRepository{
		database:   db,
		collection: collection,
	}
}

func (br *blogRepository) GetBlogs(c context.Context, filter domain.BlogRequest) ([]domain.BlogsResponse, error) {
	var result []domain.BlogsResponse
	var blogs []blog.BlogPost
	err := error(nil)

	idHex, err := bson.ObjectIDFromHex(filter.ID)
	if err == nil {
		var blogg blog.BlogPost

		collection := br.database.Collection(blog.CollectionBlogPost)
		err = collection.FindOne(c, bson.M{"_id": idHex}).Decode(&blogg)
		if err != nil {
			return result, err
		}

		item, err := PrepareBlog(br, c, blogg, filter.Content, filter.Lang)
		if err != nil {
			return result, err
		}

		result = append(result, domain.BlogsResponse{Blogs: []domain.BlogResponse{item}})
		return result, err
	}

	if filter.Limit == 0 {
		filter.Limit = 20
	}

	sortOrder := 1
	if filter.Sort == "desc" {
		sortOrder = -1
	}

	query := bson.M{}

	if filter.AllowComments {
		query["allow_comments"] = filter.AllowComments
	}

	limit := int64(filter.Limit)

	for _, value := range filter.Filters {
		// "contains", "eq", etc.
		if value.Operator == "contains" {
			query[value.Field] = bson.M{"$regex": value.Value, "$options": "i"}
		} else if value.Operator == "not_contains" {
			query[value.Field] = bson.M{"$not": bson.M{"$regex": value.Value, "$options": "i"}}
		} else {
			query[value.Field] = value.Value
		}

	}

	findOptions := options.Find().
		SetSort(bson.D{{Key: "_id", Value: sortOrder}}).
		SetLimit(limit)

	collection := br.database.Collection(blog.CollectionBlogPost)
	cursor, err := collection.Find(c, query, findOptions)
	if err != nil {
		return result, err
	}

	err = cursor.All(c, &blogs)
	if err != nil {
		return result, err
	}

	var items []domain.BlogResponse
	for i := range blogs {
		item, err := PrepareBlog(br, c, blogs[i], filter.Content, filter.Lang)
		if err != nil {
			return result, err
		}
		items = append(items, item)
	}

	result = append(result, domain.BlogsResponse{Blogs: items})

	return result, err
}

func PrepareBlog(br *blogRepository, c context.Context, blog blog.BlogPost, content []string, lang string) (domain.BlogResponse, error) {
	var result domain.BlogResponse
	err := error(nil)

	for i := range content {
		switch content[i] {
		case "comments":
			result.Coments, err = PrepareBlogComments(br, c, blog)
		}
	}

	if lang != "" {
		result.Blog, err = PrepareBlogLang(br, c, blog, lang)
	} else {
		result.Blog = blog

	}

	return result, err
}

func PrepareBlogComments(tr *blogRepository, c context.Context, item blog.BlogPost) ([]blog.BlogComment, error) {
	var comments []blog.BlogComment
	collection := tr.database.Collection(blog.CollectionBlogComment)
	cursor, err := collection.Find(c, bson.M{"blog_post_id": item.ID})
	if err != nil {
		return comments, err
	}
	err = cursor.All(c, &comments)
	if err != nil {
		return comments, err
	}
	return comments, err
}

func PrepareBlogLang(br *blogRepository, c context.Context, item blog.BlogPost, lang string) (blog.BlogPost, error) {
	var itemLang = item

	locale, err := GetLangugaBlogByCode(br, c, lang)
	if err != nil {
		return itemLang, err
	}

	record, err := GetRecordBlogByCode(br, c, blog.CollectionBlogPost)
	if err != nil {
		return itemLang, err
	}

	var items []localization.LocalizedProperty
	collection := br.database.Collection(localization.CollectionLocalizedProperty)
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
		case "body":
			itemLang.Body = items[i].LocaleValue
		case "body_overview":
			itemLang.BodyOverview = items[i].LocaleValue
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

func GetLangugaBlogByCode(tr *blogRepository, c context.Context, lang string) (localization.Language, error) {
	collection := tr.database.Collection(localization.CollectionLanguage)
	var item localization.Language
	err := collection.FindOne(c, bson.M{"unique_seo_code": lang}).Decode(&item)
	if err != nil {
		return item, err
	}
	return item, err
}

func GetRecordBlogByCode(tr *blogRepository, c context.Context, name string) (security.PermissionRecord, error) {
	collection := tr.database.Collection(security.CollectionPermissionRecord)
	var item security.PermissionRecord
	err := collection.FindOne(c, bson.M{"system_name": name}).Decode(&item)
	if err != nil {
		return item, err
	}
	return item, err
}
