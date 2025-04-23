package repository

import (
	"context"
	gdpr "earnforglance/server/domain/gdpr"
	localization "earnforglance/server/domain/localization"
	domain "earnforglance/server/domain/public"
	security "earnforglance/server/domain/security"
	"earnforglance/server/service/data/mongo"

	"go.mongodb.org/mongo-driver/v2/bson"

	"strconv"

	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type gdprConsentRepository struct {
	database   mongo.Database
	collection string
}

func NewGdprConsentRepository(db mongo.Database, collection string) domain.GdprConsentRepository {
	return &gdprConsentRepository{
		database:   db,
		collection: collection,
	}
}

func (cr *gdprConsentRepository) GetGdprConsents(c context.Context, filter domain.GdprConsentRequest) ([]domain.GdprConsentsResponse, error) {
	var result []domain.GdprConsentsResponse
	var gdprConsents []gdpr.GdprConsent

	idHex, err := bson.ObjectIDFromHex(filter.ID)
	if err == nil {
		var gdprConsentRecord gdpr.GdprConsent

		collection := cr.database.Collection(gdpr.CollectionGdprConsent)
		err = collection.FindOne(c, bson.M{"_id": idHex}).Decode(&gdprConsentRecord)
		if err != nil {
			return result, err
		}

		item, err := PrepareGdprConsent(cr, c, gdprConsentRecord, filter.Content, filter.Lang)
		if err != nil {
			return result, err
		}

		result = append(result, domain.GdprConsentsResponse{GdprConsents: []domain.GdprConsentResponse{item}})
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

	limit := int64(filter.Limit)
	//skip := int64(filter.Page * filter.Limit)

	if filter.IsRequired {
		filter.Filters = append(filter.Filters, domain.Filter{Field: "is_required", Operator: "eq", Value: strconv.FormatBool(filter.IsRequired)})
	}
	if filter.DisplayDuringRegistration {
		filter.Filters = append(filter.Filters, domain.Filter{Field: "display_during_registration", Operator: "eq", Value: strconv.FormatBool(filter.DisplayDuringRegistration)})
	}
	if filter.DisplayOnCustomerInfoPage {
		filter.Filters = append(filter.Filters, domain.Filter{Field: "display_on_customer_info_page", Operator: "eq", Value: strconv.FormatBool(filter.DisplayOnCustomerInfoPage)})
	}

	for _, value := range filter.Filters {
		// "contains", "eq", etc.
		if value.Operator == "contains" {
			query[value.Field] = bson.M{"$regex": value.Value, "$options": "i"}
		} else if value.Operator == "not_contains" {
			query[value.Field] = bson.M{"$not": bson.M{"$regex": value.Value, "$options": "i"}}
		} else {
			query[value.Field] = value.Value
		}

		//skip = 0
	}

	findOptions := options.Find().
		SetSort(bson.D{{Key: "display_order", Value: sortOrder}}).
		SetLimit(limit)

	collection := cr.database.Collection(gdpr.CollectionGdprConsent)
	cursor, err := collection.Find(c, query, findOptions)
	if err != nil {
		return result, err
	}

	err = cursor.All(c, &gdprConsents)
	if err != nil {
		return result, err
	}

	var items []domain.GdprConsentResponse
	for i := range gdprConsents {
		item, err := PrepareGdprConsent(cr, c, gdprConsents[i], filter.Content, filter.Lang)
		if err != nil {
			return result, err
		}
		items = append(items, item)
	}

	result = append(result, domain.GdprConsentsResponse{GdprConsents: items})

	return result, err
}

func PrepareGdprConsent(vr *gdprConsentRepository, c context.Context, gdprConsent gdpr.GdprConsent, content []string, lang string) (domain.GdprConsentResponse, error) {
	var result domain.GdprConsentResponse
	err := error(nil)

	for i := range content {
		switch content[i] {
		case "parent":
			parent := gdprConsent.ParentID
			if parent != nil {
				result.Parent, err = PrepareGdprParent(vr, c, *parent)
			}
		}
	}

	if lang != "" {
		result.GdprConsent, err = PrepareGdprConsentLang(vr, c, gdprConsent, lang)
	} else {
		result.GdprConsent = gdprConsent
	}

	return result, err
}

func PrepareGdprConsentLang(vr *gdprConsentRepository, c context.Context, manufacturer gdpr.GdprConsent, lang string) (gdpr.GdprConsent, error) {
	var gdprLang = manufacturer
	err := error(nil)
	locale, err := GetGdprConsentLangugaByCode(vr, c, lang)
	if err != nil {
		return gdprLang, err
	}

	record, err := GetGdprConsentRecordByCode(vr, c, gdpr.CollectionGdprConsent)
	if err != nil {
		return gdprLang, err
	}

	var items []localization.LocalizedProperty
	collection := vr.database.Collection(localization.CollectionLocalizedProperty)
	cursor, err := collection.Find(c, bson.M{"entity_id": record.ID, "language_id": locale.ID, "locale_key_group": manufacturer.ID.Hex()})

	if err != nil {
		return gdprLang, err
	}

	err = cursor.All(c, &items)
	if err != nil {
		return gdprLang, err
	}

	for i := range items {
		switch items[i].LocaleKey {
		case "title":
			gdprLang.Title = items[i].LocaleValue
		case "message":
			gdprLang.Message = items[i].LocaleValue
		case "required_message":
			gdprLang.RequiredMessage = items[i].LocaleValue
		}
	}

	return gdprLang, err
}

func GetGdprConsentLangugaByCode(vr *gdprConsentRepository, c context.Context, lang string) (localization.Language, error) {
	collection := vr.database.Collection(localization.CollectionLanguage)
	var item localization.Language
	err := collection.FindOne(c, bson.M{"unique_seo_code": lang}).Decode(&item)
	return item, err
}

func GetGdprConsentRecordByCode(vr *gdprConsentRepository, c context.Context, name string) (security.PermissionRecord, error) {
	collection := vr.database.Collection(security.CollectionPermissionRecord)
	var item security.PermissionRecord
	err := collection.FindOne(c, bson.M{"system_name": name}).Decode(&item)
	return item, err
}

func PrepareGdprParent(vr *gdprConsentRepository, c context.Context, ID bson.ObjectID) (*gdpr.GdprConsent, error) {
	var gdprConsent gdpr.GdprConsent
	collection := vr.database.Collection(gdpr.CollectionGdprConsent)
	err := collection.FindOne(c, bson.M{"_id": ID}).Decode(&gdprConsent)
	return &gdprConsent, err
}
