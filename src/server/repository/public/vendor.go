package repository

import (
	"context"
	common "earnforglance/server/domain/common"
	localization "earnforglance/server/domain/localization"
	media "earnforglance/server/domain/media"
	domain "earnforglance/server/domain/public"
	security "earnforglance/server/domain/security"
	vendor "earnforglance/server/domain/vendors"
	"earnforglance/server/service/data/mongo"
	"strconv"

	"go.mongodb.org/mongo-driver/v2/bson"

	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type vendorRepository struct {
	database   mongo.Database
	collection string
}

func NewVendorRepository(db mongo.Database, collection string) domain.VendorRepository {
	return &vendorRepository{
		database:   db,
		collection: collection,
	}
}

func (cr *vendorRepository) GetVendors(c context.Context, filter domain.VendorRequest) ([]domain.VendorsResponse, error) {
	var result []domain.VendorsResponse
	var vendors []vendor.Vendor

	idHex, err := bson.ObjectIDFromHex(filter.ID)
	if err == nil {
		var vendorRecord vendor.Vendor

		collection := cr.database.Collection(vendor.CollectionVendor)
		err = collection.FindOne(c, bson.M{"_id": idHex, "deleted": false, "active": true}).Decode(&vendorRecord)
		if err != nil {
			return result, err
		}

		item, err := PrepareVendor(cr, c, vendorRecord, filter.Content, filter.Lang)
		if err != nil {
			return result, err
		}

		result = append(result, domain.VendorsResponse{Vendors: []domain.VendorResponse{item}})
		return result, err
	}

	if filter.Limit == 0 {
		filter.Limit = 20
	}

	sortOrder := 1
	if filter.Sort == "desc" {
		sortOrder = -1
	}

	query := bson.M{"deleted": false, "active": true}

	if filter.PriceRangeFiltering {
		query["price_range_filtering"] = filter.PriceRangeFiltering
	}

	if filter.ManuallyPriceRange {
		query["manually_price_range"] = filter.ManuallyPriceRange
	}

	if filter.PriceFrom > 0 {
		query["price_from"] = bson.M{"$gte": filter.PriceFrom}
	}

	if filter.PriceTo > 0 {
		query["price_to"] = bson.M{"$lte": filter.PriceTo}
	}

	limit := int64(filter.Limit)
	//skip := int64(filter.Page * filter.Limit)

	if filter.ManuallyPriceRange {
		filter.Filters = append(filter.Filters, domain.Filter{Field: "is_required", Operator: "eq", Value: strconv.FormatBool(filter.ManuallyPriceRange)})
	}
	if filter.PriceRangeFiltering {
		filter.Filters = append(filter.Filters, domain.Filter{Field: "display_during_registration", Operator: "eq", Value: strconv.FormatBool(filter.PriceRangeFiltering)})
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

	collection := cr.database.Collection(vendor.CollectionVendor)
	cursor, err := collection.Find(c, query, findOptions)
	if err != nil {
		return result, err
	}

	err = cursor.All(c, &vendors)
	if err != nil {
		return result, err
	}

	var items []domain.VendorResponse
	for i := range vendors {
		item, err := PrepareVendor(cr, c, vendors[i], filter.Content, filter.Lang)
		if err != nil {
			return result, err
		}
		items = append(items, item)
	}

	result = append(result, domain.VendorsResponse{Vendors: items})

	return result, err
}

func PrepareVendor(vr *vendorRepository, c context.Context, vendor vendor.Vendor, content []string, lang string) (domain.VendorResponse, error) {
	var result domain.VendorResponse
	err := error(nil)

	for i := range content {
		switch content[i] {
		case "address":
			result.Address, err = PrepareVendorAddress(vr, c, vendor.AddressID)
		case "picture":
			result.Picture, err = PrepareVendorPicture(vr, c, vendor.PictureID)
		}
	}

	if lang != "" {
		result.Vendor, err = PrepareVendorLang(vr, c, vendor, lang)
	} else {
		result.Vendor = vendor
	}

	return result, err
}

func PrepareVendorLang(vr *vendorRepository, c context.Context, manufacturer vendor.Vendor, lang string) (vendor.Vendor, error) {
	var manufacturerLang = manufacturer
	err := error(nil)
	locale, err := GetVendorLangugaByCode(vr, c, lang)
	if err != nil {
		return manufacturerLang, err
	}

	record, err := GetVendorRecordByCode(vr, c, vendor.CollectionVendor)
	if err != nil {
		return manufacturerLang, err
	}

	var items []localization.LocalizedProperty
	collection := vr.database.Collection(localization.CollectionLocalizedProperty)
	cursor, err := collection.Find(c, bson.M{"entity_id": record.ID, "language_id": locale.ID, "locale_key_group": manufacturer.ID.Hex()})

	if err != nil {
		return manufacturerLang, err
	}

	err = cursor.All(c, &items)
	if err != nil {
		return manufacturerLang, err
	}

	for i := range items {
		switch items[i].LocaleKey {
		case "name":
			manufacturerLang.Name = items[i].LocaleValue
		case "description":
			manufacturerLang.Description = items[i].LocaleValue
		case "meta_title":
			manufacturerLang.MetaTitle = items[i].LocaleValue
		case "meta_keywords":
			manufacturerLang.MetaKeywords = items[i].LocaleValue
		case "meta_description":
			manufacturerLang.MetaDescription = items[i].LocaleValue
		}
	}

	return manufacturerLang, err
}

func GetVendorLangugaByCode(vr *vendorRepository, c context.Context, lang string) (localization.Language, error) {
	collection := vr.database.Collection(localization.CollectionLanguage)
	var item localization.Language
	err := collection.FindOne(c, bson.M{"unique_seo_code": lang}).Decode(&item)
	return item, err
}

func GetVendorRecordByCode(vr *vendorRepository, c context.Context, name string) (security.PermissionRecord, error) {
	collection := vr.database.Collection(security.CollectionPermissionRecord)
	var item security.PermissionRecord
	err := collection.FindOne(c, bson.M{"system_name": name}).Decode(&item)
	return item, err
}

func PrepareVendorPicture(vr *vendorRepository, c context.Context, ID bson.ObjectID) (*media.Picture, error) {
	var picture media.Picture
	collection := vr.database.Collection(media.CollectionPicture)
	err := collection.FindOne(c, bson.M{"_id": ID}).Decode(&picture)
	return &picture, err
}

func PrepareVendorAddress(vr *vendorRepository, c context.Context, ID bson.ObjectID) (common.Address, error) {
	var address common.Address
	collection := vr.database.Collection(common.CollectionAddress)
	err := collection.FindOne(c, bson.M{"_id": ID}).Decode(&address)
	return address, err
}
