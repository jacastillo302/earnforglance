package repository

import (
	"context"
	catalog "earnforglance/server/domain/catalog"
	customer "earnforglance/server/domain/customers"
	localization "earnforglance/server/domain/localization"
	media "earnforglance/server/domain/media"
	domain "earnforglance/server/domain/public"
	security "earnforglance/server/domain/security"
	shipping "earnforglance/server/domain/shipping"
	tax "earnforglance/server/domain/tax"
	vendor "earnforglance/server/domain/vendors"
	"earnforglance/server/service/data/mongo"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type catalogRepository struct {
	database   mongo.Database
	collection string
}

func NewCatalogRepository(db mongo.Database, collection string) domain.CatalogRepository {
	return &catalogRepository{
		database:   db,
		collection: collection,
	}
}

func (cr *catalogRepository) GetCategories(c context.Context, filter domain.CategoryRequest) ([]domain.CategoriesResponse, error) {
	var result []domain.CategoriesResponse
	var categories []catalog.Category

	idHex, err := primitive.ObjectIDFromHex(filter.ID)
	if err == nil {
		var category catalog.Category

		collection := cr.database.Collection(catalog.CollectionCategory)
		err = collection.FindOne(c, bson.M{"_id": idHex, "deleted": false, "published": true}).Decode(&category)
		if err != nil {
			return result, err
		}

		item, err := PrepareCategory(cr, c, category, filter.Content, filter.Lang)
		if err != nil {
			return result, err
		}

		result = append(result, domain.CategoriesResponse{Categories: []domain.CategoryResponse{item}})
		return result, err
	}

	if filter.Limit == 0 {
		filter.Limit = 20
	}

	sortOrder := 1
	if filter.Sort == "desc" {
		sortOrder = -1
	}

	// Build dynamic filter
	query := bson.M{"deleted": false}

	if filter.ShowOnHomepage {
		query["show_on_homepage"] = filter.ShowOnHomepage
	}

	if filter.IncludeInTopMenu {
		query["include_in_top_menu"] = filter.IncludeInTopMenu
	}

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

	if filter.Parent != "" {
		idHex, err := primitive.ObjectIDFromHex(filter.Parent)
		if err == nil {
			query["parent_category_id"] = idHex
		}
	}

	limit := int64(filter.Limit)
	skip := int64(filter.Page * filter.Limit)

	for _, value := range filter.Filters {
		// "contains", "eq", etc.
		if value.Operator == "contains" {
			query[value.Field] = bson.M{"$regex": value.Value, "$options": "i"}
		} else {
			query[value.Field] = value.Value
		}

		skip = 0
	}

	findOptions := options.Find().
		SetSort(bson.D{{Key: "_id", Value: sortOrder}}).
		SetLimit(int64(limit)).
		SetSkip(skip)

	collection := cr.database.Collection(catalog.CollectionCategory)
	cursor, err := collection.Find(c, query, findOptions)
	if err != nil {
		return result, err
	}

	err = cursor.All(c, &categories)
	if err != nil {
		return result, err
	}

	var items []domain.CategoryResponse
	for i := range categories {
		item, err := PrepareCategory(cr, c, categories[i], filter.Content, filter.Lang)
		if err != nil {
			return result, err
		}
		items = append(items, item)
	}

	result = append(result, domain.CategoriesResponse{Categories: items})

	return result, err
}

func PrepareCategory(cr *catalogRepository, c context.Context, category catalog.Category, content []string, lang string) (domain.CategoryResponse, error) {
	var result domain.CategoryResponse

	for i := range content {
		switch content[i] {
		case "template":
			result.Template, _ = PrepareCategoryTemplate(cr, c, category)
		case "picture":
			result.Picture, _ = PrepareCategoryPicture(cr, c, category)
		case "childs":
			result.Childs, _ = PrepareCategoryChilds(cr, c, category)
		}
	}

	if lang != "" {
		result.Category, _ = PrepareCategoryLang(cr, c, category, lang)
	} else {
		result.Category = category
	}

	return result, nil
}

func (cr *catalogRepository) GetProducts(c context.Context, filter domain.ProductRequest) ([]domain.ProductsResponse, error) {
	var result []domain.ProductsResponse
	var products []catalog.Product

	idHex, err := primitive.ObjectIDFromHex(filter.ID)
	if err == nil {
		var product catalog.Product

		collection := cr.database.Collection(catalog.CollectionProduct)
		err = collection.FindOne(c, bson.M{"_id": idHex, "deleted": false, "published": true}).Decode(&product)
		if err != nil {
			return result, err
		}

		item, err := PrepareProduct(cr, c, product, filter.Content, filter.Lang)
		if err != nil {
			return result, err
		}

		result = append(result, domain.ProductsResponse{Products: []domain.ProductResponse{item}})
		return result, err
	}

	if filter.Limit == 0 {
		filter.Limit = 20
	}

	sortOrder := 1
	if filter.Sort == "desc" {
		sortOrder = -1
	}

	// Build dynamic filter
	query := bson.M{"deleted": false}

	if filter.ShowOnHomepage {
		query["show_on_homepage"] = filter.ShowOnHomepage
	}

	if filter.IsRental {
		query["is_rental"] = filter.IsRental
	}

	if filter.IsTaxExempt {
		query["is_tax_exempt"] = filter.IsTaxExempt
	}

	if filter.MarkAsNew {
		query["mark_as_new"] = filter.MarkAsNew
	}

	if filter.MinPrice > 0 && filter.MxnPrice > 0 {
		query["price"] = bson.M{
			"$gte": filter.MinPrice,
			"$lte": filter.MxnPrice,
		}
	} else if filter.MinPrice > 0 {
		query["price"] = bson.M{
			"$gte": filter.MinPrice,
		}
	} else if filter.MxnPrice > 0 {
		query["price"] = bson.M{
			"$lte": filter.MxnPrice,
		}
	}

	if len(filter.Categories) > 0 {
		var products []primitive.ObjectID
		var categories []catalog.ProductCategory
		for i := range filter.Categories {
			idHex, err := primitive.ObjectIDFromHex(filter.Categories[i])
			if err == nil {
				filter.Categories[i] = idHex.Hex()
				collection := cr.database.Collection(catalog.CollectionProductCategory)
				cursor, err := collection.Find(c, bson.M{"_id": idHex, "deleted": false})
				if err != nil {
					return result, err
				}

				err = cursor.All(c, &categories)
				if err != nil {
					return result, err
				}

				for i := range categories {
					products = append(products, categories[i].ProductID)
				}
			}
		}

		if len(products) > 0 {
			query["product_id"] = bson.M{"$in": products}
		} else {
			query["product_id"] = nil
		}
	}

	limit := int64(filter.Limit)
	skip := int64(filter.Page * filter.Limit)

	for _, value := range filter.Filters {
		// "contains", "eq", etc.
		if value.Operator == "contains" {
			query[value.Field] = bson.M{"$regex": value.Value, "$options": "i"}
		} else {
			query[value.Field] = value.Value
		}

		skip = 0
	}

	findOptions := options.Find().
		SetSort(bson.D{{Key: "_id", Value: sortOrder}}).
		SetLimit(int64(limit)).
		SetSkip(skip)

	collection := cr.database.Collection(catalog.CollectionProduct)
	cursor, err := collection.Find(c, query, findOptions)
	if err != nil {
		return result, err
	}

	err = cursor.All(c, &products)
	if err != nil {
		return result, err
	}

	var items []domain.ProductResponse
	for i := range products {
		item, err := PrepareProduct(cr, c, products[i], filter.Content, filter.Lang)
		if err != nil {
			return result, err
		}

		items = append(items, item)
	}
	result = append(result, domain.ProductsResponse{Products: items})

	return result, err
}

func PrepareProduct(cr *catalogRepository, c context.Context, product catalog.Product, content []string, lang string) (domain.ProductResponse, error) {
	var result domain.ProductResponse

	for i := range content {
		switch content[i] {
		case "template":
			result.Template, _ = PrepareProductTemplate(cr, c, product)
		case "categories":
			result.Categories, _ = PrepareProductCategory(cr, c, product)
		case "specifications":
			result.Specifications, _ = PrepareProductSpecificationAttribute(cr, c, product)
		case "attributes":
			result.Attributes, _ = PrepareProductAttribute(cr, c, product)
		case "warehouse":
			result.Warehouse.Warehouse, _ = PrepareProductWarehouse(cr, c, product)
			result.Warehouse.Inventory, _ = PrepareProductWarehouseInventory(cr, c, product)
		case "delivery":
			result.DeliveryDate, _ = PrepareProductDeliveryDate(cr, c, product)
		case "range":
			result.Range, _ = PrepareProductAvailabilityRange(cr, c, product)
		case "tax":
			result.Tax, _ = PrepareTaxCategory(cr, c, product.TaxCategoryID)
		case "vendor":
			result.Vendor, _ = PrepareVendor(cr, c, product.VendorID)
		case "reviews":
			result.Reviews, _ = PrepareProductReview(cr, c, product)
		case "download":
			download, _ := PrepareDownload(cr, c, product.DownloadID)
			result.Download = &download
		case "tierprices":
			result.TierPrice, _ = PrepareProductTierPrice(cr, c, product)
		case "cross":
			result.Cross, _ = PrepareCrossSellProduct(cr, c, product)
		case "relates":
			result.Relates, _ = PrepareRelatedProduct(cr, c, product)
		case "tags":
			result.Tags, _ = PrepareProductTag(cr, c, product)
		case "manufacturers":
			result.Manufacturers, _ = PrepareProductManufacturer(cr, c, product)
		case "videos":
			result.Videos, _ = PrepareProductVideo(cr, c, product)
		case "pictures":
			result.Pictures, _ = PrepareProductPicture(cr, c, product)
		}
	}

	if lang != "" {
		result.Product, _ = PrepareProductLang(cr, c, product, lang)
	} else {
		result.Product = product
	}

	return result, nil
}

func GetLangugaByCode(cr *catalogRepository, c context.Context, lang string) (localization.Language, error) {
	collection := cr.database.Collection(localization.CollectionLanguage)
	var item localization.Language
	err := collection.FindOne(c, bson.M{"unique_seo_code": lang}).Decode(&item)
	return item, err
}

func GetRecordByCode(cr *catalogRepository, c context.Context, name string) (security.PermissionRecord, error) {
	collection := cr.database.Collection(security.CollectionPermissionRecord)
	var item security.PermissionRecord
	err := collection.FindOne(c, bson.M{"system_name": name}).Decode(&item)
	return item, err
}

func PrepareCategoryLang(cr *catalogRepository, c context.Context, category catalog.Category, lang string) (catalog.Category, error) {
	var categoryLang = category

	locale, err := GetLangugaByCode(cr, c, lang)
	if err != nil {
		return categoryLang, err
	}

	record, err := GetRecordByCode(cr, c, catalog.CollectionCategory)
	if err != nil {
		return categoryLang, err
	}

	var items []localization.LocalizedProperty
	collection := cr.database.Collection(localization.CollectionLocalizedProperty)
	cursor, err := collection.Find(c, bson.M{"entity_id": record.ID, "language_id": locale.ID, "locale_key_group": category.ID.Hex()})

	if err != nil {
		return categoryLang, err
	}

	err = cursor.All(c, &items)
	if err != nil {
		return categoryLang, err
	}

	for i := range items {
		switch items[i].LocaleKey {
		case "name":
			categoryLang.Name = items[i].LocaleValue
		case "description":
			categoryLang.Description = items[i].LocaleValue
		case "meta_title":
			categoryLang.MetaTitle = items[i].LocaleValue
		case "meta_keywords":
			categoryLang.MetaKeywords = items[i].LocaleValue
		case "meta_description":
			categoryLang.MetaDescription = items[i].LocaleValue
		}
	}

	return categoryLang, nil
}

func PrepareProductLang(cr *catalogRepository, c context.Context, product catalog.Product, lang string) (catalog.Product, error) {
	var productLang = product

	locale, err := GetLangugaByCode(cr, c, lang)
	if err != nil {
		return productLang, err
	}

	record, err := GetRecordByCode(cr, c, catalog.CollectionProduct)
	if err != nil {
		return productLang, err
	}

	var items []localization.LocalizedProperty
	collection := cr.database.Collection(localization.CollectionLocalizedProperty)
	cursor, err := collection.Find(c, bson.M{"entity_id": record.ID, "language_id": locale.ID, "locale_key_group": product.ID.Hex()})
	if err != nil {
		return productLang, err
	}

	err = cursor.All(c, &items)
	if err != nil {
		return productLang, err
	}

	for i := range items {
		switch items[i].LocaleKey {
		case "name":
			productLang.Name = items[i].LocaleValue
		case "full_description":
			productLang.FullDescription = items[i].LocaleValue
		case "short_description":
			productLang.ShortDescription = items[i].LocaleValue
		case "meta_title":
			productLang.MetaTitle = items[i].LocaleValue
		case "meta_keywords":
			productLang.MetaKeywords = items[i].LocaleValue
		case "meta_description":
			productLang.MetaDescription = items[i].LocaleValue
		}
	}

	return productLang, nil
}

func PrepareProductTemplate(cr *catalogRepository, c context.Context, product catalog.Product) (catalog.ProductTemplate, error) {

	var template catalog.ProductTemplate
	collection := cr.database.Collection(catalog.CollectionProductTemplate)
	err := collection.FindOne(c, bson.M{"_id": product.ProductTemplateID}).Decode(&template)

	return template, err

}

func PrepareCategoryTemplate(cr *catalogRepository, c context.Context, category catalog.Category) (catalog.CategoryTemplate, error) {

	var template catalog.CategoryTemplate
	collection := cr.database.Collection(catalog.CollectionCategoryTemplate)
	err := collection.FindOne(c, bson.M{"_id": category.CategoryTemplateID}).Decode(&template)

	return template, err

}

func PrepareProductReview(cr *catalogRepository, c context.Context, product catalog.Product) ([]domain.ProductReview, error) {

	var reviews []domain.ProductReview

	var productreview []catalog.ProductReview
	collection := cr.database.Collection(catalog.CollectionProductReview)
	cursor, err := collection.Find(c, bson.M{"product_id": product.ID})
	if err != nil {
		return nil, err
	}

	err = cursor.All(c, &productreview)
	if err != nil {
		return nil, err
	}

	for i := range productreview {

		var reviewTypeMap catalog.ProductReviewReviewTypeMapping
		collection = cr.database.Collection(catalog.CollectionProductReviewReviewTypeMapping)
		err = collection.FindOne(c, bson.M{"product_review_id": productreview[i].ID}).Decode(&reviewTypeMap)
		var reviewType catalog.ReviewType
		if err == nil {
			collection = cr.database.Collection(catalog.CollectionReviewType)
			collection.FindOne(c, bson.M{"_id": reviewTypeMap.ReviewTypeID}).Decode(&reviewType)
		}

		var reviewHelps []catalog.ProductReviewHelpfulness
		collection = cr.database.Collection(catalog.CollectionProductReviewHelpfulness)
		cursor, err := collection.Find(c, bson.M{"product_review_id": productreview[i].ID})
		if err != nil {
			return nil, err
		}

		err = cursor.All(c, &reviewHelps)
		if err != nil {
			return nil, err
		}

		var customerReplay customer.Customer
		collection = cr.database.Collection(customer.CollectionCustomer)
		collection.FindOne(c, bson.M{"_id": reviewTypeMap.ReviewTypeID}).Decode(&customerReplay)

		reviews = append(reviews, domain.ProductReview{Review: productreview[0], Type: reviewType.Name, Customer: customerReplay.FirstName, Helpfulness: reviewHelps})
	}

	return reviews, err
}

func PrepareProductCategory(cr *catalogRepository, c context.Context, product catalog.Product) ([]catalog.Category, error) {

	var productcategory []catalog.ProductCategory
	collection := cr.database.Collection(catalog.CollectionProductCategory)
	cursor, err := collection.Find(c, bson.M{"product_id": product.ID})
	if err != nil {
		return nil, err
	}

	err = cursor.All(c, &productcategory)
	if err != nil {
		return nil, err
	}

	var categories []catalog.Category
	var category catalog.Category
	collection = cr.database.Collection(catalog.CollectionCategory)
	for i := range productcategory {
		err = collection.FindOne(c, bson.M{"_id": productcategory[i].CategoryID}).Decode(&category)
		if err == nil {
			categories = append(categories, category)
		}
	}

	return categories, err
}

func PrepareCategoryChilds(cr *catalogRepository, c context.Context, category catalog.Category) ([]domain.CategoryChilds, error) {

	var result []domain.CategoryChilds
	var categories []catalog.Category

	collection := cr.database.Collection(catalog.CollectionCategory)
	cursor, err := collection.Find(c, bson.M{"parent_category_id": category.ID})
	if err != nil {
		return nil, err
	}

	err = cursor.All(c, &categories)
	if err != nil {
		return nil, err
	}

	for i := range categories {
		picture, _ := PrepareCategoryPicture(cr, c, category)
		result = append(result, domain.CategoryChilds{Category: categories[i], Picture: picture})
	}

	return result, err
}

func PrepareProductSpecificationAttribute(cr *catalogRepository, c context.Context, product catalog.Product) ([]domain.SpecificationAttribute, error) {
	specificationAttributes := []domain.SpecificationAttribute{}

	var productSpecificationAttribute []catalog.ProductSpecificationAttribute
	collection := cr.database.Collection(catalog.CollectionProductSpecificationAttribute)
	cursor, err := collection.Find(c, bson.M{"product_id": product.ID})
	if err != nil {
		return specificationAttributes, err
	}

	err = cursor.All(c, &productSpecificationAttribute)
	if err != nil {
		return specificationAttributes, err
	}

	var specificationAttribute catalog.SpecificationAttribute
	var specificationAttributeGroup catalog.SpecificationAttributeGroup
	var specificationAttributeOptions []catalog.SpecificationAttributeOption

	bNewAttribute := false
	bSaveAttribute := false
	var idTemp = ""
	for i := range productSpecificationAttribute {
		var specificationAttributeTemp catalog.SpecificationAttribute
		var specificationAttributeGroupTemp catalog.SpecificationAttributeGroup
		var specificationAttributeOption catalog.SpecificationAttributeOption

		collection = cr.database.Collection(catalog.CollectionSpecificationAttributeOption)
		collection.FindOne(c, bson.M{"_id": productSpecificationAttribute[i].SpecificationAttributeOptionID}).Decode(&specificationAttributeOption)

		collection = cr.database.Collection(catalog.CollectionSpecificationAttribute)
		collection.FindOne(c, bson.M{"_id": specificationAttributeOption.SpecificationAttributeID}).Decode(&specificationAttributeTemp)

		collection = cr.database.Collection(catalog.CollectionSpecificationAttributeGroup)
		collection.FindOne(c, bson.M{"_id": specificationAttributeTemp.SpecificationAttributeGroupID}).Decode(&specificationAttributeGroupTemp)

		if idTemp == "" {
			idTemp = specificationAttributeTemp.ID.Hex()
			specificationAttributeOptions = append(specificationAttributeOptions, specificationAttributeOption)
			bSaveAttribute = true
		} else {

			if idTemp != specificationAttributeTemp.ID.Hex() {
				idTemp = specificationAttributeTemp.ID.Hex()
				bNewAttribute = true
			} else {
				specificationAttributeOptions = append(specificationAttributeOptions, specificationAttributeOption)
				bSaveAttribute = true
				bNewAttribute = false
			}
		}

		if !bNewAttribute {
			specificationAttribute = specificationAttributeTemp
			specificationAttributeGroup = specificationAttributeGroupTemp
		} else {

			specificationAttributes = append(specificationAttributes, domain.SpecificationAttribute{Attribute: specificationAttribute, Options: specificationAttributeOptions, Group: specificationAttributeGroup})

			specificationAttributeOptions = nil
			specificationAttributeOptions = append(specificationAttributeOptions, specificationAttributeOption)
			specificationAttribute = specificationAttributeTemp
			specificationAttributeGroup = specificationAttributeGroupTemp

			bNewAttribute = false
			bSaveAttribute = true
		}

	}

	if bSaveAttribute {
		specificationAttributes = append(specificationAttributes, domain.SpecificationAttribute{Attribute: specificationAttribute, Options: specificationAttributeOptions, Group: specificationAttributeGroup})
		specificationAttributeOptions = nil
	}

	return specificationAttributes, err
}

func PrepareProductAttribute(cr *catalogRepository, c context.Context, product catalog.Product) ([]domain.ProductAttribute, error) {

	var productAttributeResponse []domain.ProductAttribute
	var productAttributeMap []catalog.ProductAttributeMapping
	collection := cr.database.Collection(catalog.CollectionProductAttributeMapping)
	cursor, err := collection.Find(c, bson.M{"product_id": product.ID})
	if err != nil {
		return productAttributeResponse, err
	}

	err = cursor.All(c, &productAttributeMap)
	if err != nil {
		return productAttributeResponse, err
	}

	for i := range productAttributeMap {

		var productAttribute catalog.ProductAttribute
		collection = cr.database.Collection(catalog.CollectionProductAttribute)
		collection.FindOne(c, bson.M{"_id": productAttributeMap[i].ProductAttributeID}).Decode(&productAttribute)

		var productAttributeValues []catalog.ProductAttributeValue
		collection = cr.database.Collection(catalog.CollectionProductAttributeValue)
		cursor, err = collection.Find(c, bson.M{"product_attribute_mapping_id": productAttributeMap[i].ID})
		if err != nil {
			return nil, err
		}

		err = cursor.All(c, &productAttributeValues)
		if err != nil {
			return productAttributeResponse, err
		}

		var productAttributeValue []domain.ProductAttributeValue
		for i := range productAttributeValues {

			var valuespictures []catalog.ProductAttributeValuePicture
			collection = cr.database.Collection(catalog.CollectionProductAttributeValuePicture)
			cursor, err = collection.Find(c, bson.M{"product_attribute_value_id": productAttributeValues[i].ID})
			if err != nil {
				return productAttributeResponse, err
			}

			err = cursor.All(c, &valuespictures)
			if err != nil {
				return productAttributeResponse, err
			}

			var picturesAttribute []media.Picture
			for f := range valuespictures {
				var picture media.Picture
				collection = cr.database.Collection(media.CollectionPicture)
				collection.FindOne(c, bson.M{"_id": valuespictures[f].PictureID}).Decode(&picture)
				picturesAttribute = append(picturesAttribute, picture)
			}

			value := domain.ProductAttributeValue{Value: productAttributeValues[i], Pictures: picturesAttribute}
			productAttributeValue = append(productAttributeValue, value)

		}

		var productAttributeCombinations []domain.ProductAttributeCombination
		var catalogAttributeCombinations []catalog.ProductAttributeCombination

		collection = cr.database.Collection(catalog.CollectionProductAttributeCombination)
		cursor, err = collection.Find(c, bson.M{"product_id": product.ID})
		if err != nil {
			return productAttributeResponse, err
		}

		err = cursor.All(c, &catalogAttributeCombinations)
		if err != nil {
			return productAttributeResponse, err
		}

		for f := range catalogAttributeCombinations {

			var catalogAttributeCombinationPictures []catalog.ProductAttributeCombinationPicture
			collection = cr.database.Collection(catalog.CollectionProductAttributeCombinationPicture)
			cursor, err = collection.Find(c, bson.M{"product_id": catalogAttributeCombinations[f].ProductID})
			if err != nil {
				return productAttributeResponse, err
			}

			err = cursor.All(c, &catalogAttributeCombinationPictures)
			if err != nil {
				return productAttributeResponse, err
			}

			var pictures []media.Picture
			for f := range catalogAttributeCombinationPictures {
				var picture media.Picture
				collection = cr.database.Collection(media.CollectionPicture)
				collection.FindOne(c, bson.M{"_id": catalogAttributeCombinationPictures[f].PictureID}).Decode(&picture)
				pictures = append(pictures, picture)
			}

			productAttributeCombinations = append(productAttributeCombinations, domain.ProductAttributeCombination{Value: catalogAttributeCombinations[f], Pictures: pictures})

		}

		productAttributeResponse = append(productAttributeResponse, domain.ProductAttribute{Attribute: productAttribute, Values: productAttributeValue, Combinations: productAttributeCombinations})

	}

	return productAttributeResponse, err
}

func PrepareProductWarehouse(cr *catalogRepository, c context.Context, product catalog.Product) (shipping.Warehouse, error) {
	var warehouse shipping.Warehouse
	collection := cr.database.Collection(shipping.CollectionWarehouse)
	err := collection.FindOne(c, bson.M{"_id": product.WarehouseID}).Decode(&warehouse)
	return warehouse, err
}

func PrepareProductWarehouseInventory(cr *catalogRepository, c context.Context, product catalog.Product) (catalog.ProductWarehouseInventory, error) {
	var inventory catalog.ProductWarehouseInventory
	collection := cr.database.Collection(catalog.CollectionProductWarehouseInventory)
	err := collection.FindOne(c, bson.M{"_id": product.WarehouseID}).Decode(&inventory)
	return inventory, err
}

func PrepareProductDeliveryDate(cr *catalogRepository, c context.Context, product catalog.Product) (shipping.DeliveryDate, error) {
	var deliveryDate shipping.DeliveryDate
	collection := cr.database.Collection(shipping.CollectionDeliveryDate)
	err := collection.FindOne(c, bson.M{"_id": product.DeliveryDateID}).Decode(&deliveryDate)
	return deliveryDate, err
}

func PrepareProductAvailabilityRange(cr *catalogRepository, c context.Context, product catalog.Product) (shipping.ProductAvailabilityRange, error) {
	var availabilityRange shipping.ProductAvailabilityRange
	collection := cr.database.Collection(shipping.CollectionProductAvailabilityRange)
	err := collection.FindOne(c, bson.M{"_id": product.ProductAvailabilityRangeID}).Decode(&availabilityRange)
	return availabilityRange, err
}

func PrepareTaxCategory(cr *catalogRepository, c context.Context, ID primitive.ObjectID) (tax.TaxCategory, error) {
	var taxes tax.TaxCategory
	collection := cr.database.Collection(tax.CollectionTaxCategory)
	err := collection.FindOne(c, bson.M{"_id": ID}).Decode(&taxes)
	return taxes, err
}

func PrepareVendor(cr *catalogRepository, c context.Context, ID primitive.ObjectID) (vendor.Vendor, error) {
	var vendo vendor.Vendor
	collection := cr.database.Collection(vendor.CollectionVendor)
	err := collection.FindOne(c, bson.M{"_id": ID}).Decode(&vendo)
	return vendo, err
}
func PrepareDownload(cr *catalogRepository, c context.Context, ID primitive.ObjectID) (media.Download, error) {
	var download media.Download
	collection := cr.database.Collection(media.CollectionDownload)
	err := collection.FindOne(c, bson.M{"_id": ID}).Decode(&download)
	return download, err
}

func PrepareProductTierPrice(cr *catalogRepository, c context.Context, product catalog.Product) ([]catalog.TierPrice, error) {
	var tier []catalog.TierPrice
	collection := cr.database.Collection(catalog.CollectionTierPrice)
	cursor, err := collection.Find(c, bson.M{"product_id": product.ID})
	if err != nil {
		return tier, err
	}
	err = cursor.All(c, &tier)
	return tier, err
}

func PrepareCrossSellProduct(cr *catalogRepository, c context.Context, product catalog.Product) ([]catalog.Product, error) {
	var crossproduct []catalog.CrossSellProduct
	var result []catalog.Product
	collection := cr.database.Collection(catalog.CollectionCrossSellProduct)
	cursor, err := collection.Find(c, bson.M{"product_id1": product.ID})
	if err != nil {
		return result, err
	}

	err = cursor.All(c, &crossproduct)
	if err != nil {
		return result, err
	}

	var productcross1 catalog.Product
	collection = cr.database.Collection(catalog.CollectionProduct)
	for i := range crossproduct {
		err = collection.FindOne(c, bson.M{"_id": crossproduct[i].ProductID1}).Decode(&productcross1)
		if err == nil {
			result = append(result, productcross1)
		}
	}

	var crossproduct2 []catalog.CrossSellProduct
	collection = cr.database.Collection(catalog.CollectionCrossSellProduct)
	cursor, err = collection.Find(c, bson.M{"product_id2": product.ID})
	if err != nil {
		return result, err
	}

	err = cursor.All(c, &crossproduct2)
	if err != nil {
		return result, err
	}

	var productcross2 catalog.Product
	collection = cr.database.Collection(catalog.CollectionProduct)
	for i := range crossproduct2 {
		err = collection.FindOne(c, bson.M{"_id": crossproduct2[i].ProductID1}).Decode(&productcross2)
		if err == nil {
			result = append(result, productcross2)
		}
	}

	return result, err
}

func PrepareRelatedProduct(cr *catalogRepository, c context.Context, product catalog.Product) ([]catalog.Product, error) {
	var crossproduct []catalog.RelatedProduct
	var result []catalog.Product
	collection := cr.database.Collection(catalog.CollectionRelatedProduct)
	cursor, err := collection.Find(c, bson.M{"product_id1": product.ID})
	if err != nil {
		return result, err
	}

	err = cursor.All(c, &crossproduct)
	if err != nil {
		return result, err
	}

	var productcross1 catalog.Product
	collection = cr.database.Collection(catalog.CollectionProduct)
	for i := range crossproduct {
		err = collection.FindOne(c, bson.M{"_id": crossproduct[i].ProductID1}).Decode(&productcross1)
		if err == nil {
			result = append(result, productcross1)
		}
	}

	var crossproduct2 []catalog.RelatedProduct
	collection = cr.database.Collection(catalog.CollectionRelatedProduct)
	cursor, err = collection.Find(c, bson.M{"product_id2": product.ID})
	if err != nil {
		return result, err
	}

	err = cursor.All(c, &crossproduct2)
	if err != nil {
		return result, err
	}

	var productcross2 catalog.Product
	collection = cr.database.Collection(catalog.CollectionProduct)
	for i := range crossproduct2 {
		err = collection.FindOne(c, bson.M{"_id": crossproduct2[i].ProductID1}).Decode(&productcross2)
		if err == nil {
			result = append(result, productcross2)
		}
	}

	return result, err
}

func PrepareProductTag(cr *catalogRepository, c context.Context, product catalog.Product) ([]catalog.ProductTag, error) {
	var productagmap []catalog.ProductProductTagMapping
	var tag catalog.ProductTag
	var producta []catalog.ProductTag

	collection := cr.database.Collection(catalog.CollectionProductProductTagMapping)
	cursor, err := collection.Find(c, bson.M{"product_id": product.ID})
	if err != nil {
		return producta, err
	}

	err = cursor.All(c, &productagmap)
	if err != nil {
		return producta, err
	}

	collection = cr.database.Collection(catalog.CollectionProductTag)
	for i := range productagmap {
		err = collection.FindOne(c, bson.M{"_id": productagmap[i].ProductTagID}).Decode(&tag)
		if err == nil {
			producta = append(producta, tag)
		}
	}

	return producta, err
}

func PrepareProductManufacturer(cr *catalogRepository, c context.Context, product catalog.Product) ([]catalog.Manufacturer, error) {
	var manufacturerMaping []catalog.ProductManufacturer
	var manufacturers []catalog.Manufacturer
	collection := cr.database.Collection(catalog.CollectionProductManufacturer)
	cursor, err := collection.Find(c, bson.M{"product_id": product.ID})
	if err != nil {
		return manufacturers, err
	}

	err = cursor.All(c, &manufacturerMaping)
	if err != nil {
		return manufacturers, err
	}

	var manufacturer catalog.Manufacturer
	collection = cr.database.Collection(catalog.CollectionManufacturer)
	for i := range manufacturerMaping {
		err = collection.FindOne(c, bson.M{"_id": manufacturerMaping[i].ManufacturerID}).Decode(&manufacturer)
		if err == nil {
			manufacturers = append(manufacturers, manufacturer)
		}
	}

	return manufacturers, err
}

func PreparePicture(cr *catalogRepository, c context.Context, ID primitive.ObjectID) (media.Picture, error) {
	var picture media.Picture
	collection := cr.database.Collection(media.CollectionPicture)
	err := collection.FindOne(c, bson.M{"_id": ID}).Decode(&picture)
	return picture, err
}

func PrepareProductPicture(cr *catalogRepository, c context.Context, product catalog.Product) ([]media.Picture, error) {
	var productPictures []catalog.ProductPicture
	var pictures []media.Picture
	collection := cr.database.Collection(catalog.CollectionProductPicture)
	cursor, err := collection.Find(c, bson.M{"product_id": product.ID})
	if err != nil {
		return pictures, err
	}

	err = cursor.All(c, &productPictures)
	if err != nil {
		return pictures, err
	}

	var picture media.Picture
	for i := range productPictures {
		picture, err = PreparePicture(cr, c, productPictures[i].PictureID)
		if err != nil {
			return pictures, err
		}
		pictures = append(pictures, picture)
	}

	return pictures, err
}

func PrepareCategoryPicture(cr *catalogRepository, c context.Context, category catalog.Category) (media.Picture, error) {
	var picture media.Picture
	picture, err := PreparePicture(cr, c, category.PictureID)
	if err != nil {
		return picture, err
	}
	return picture, err
}

func PrepareVideo(cr *catalogRepository, c context.Context, ID primitive.ObjectID) (media.Video, error) {

	var video media.Video
	collection := cr.database.Collection(media.CollectionVideo)
	err := collection.FindOne(c, bson.M{"_id": ID}).Decode(&video)
	return video, err
}

func PrepareProductVideo(cr *catalogRepository, c context.Context, product catalog.Product) ([]media.Video, error) {
	var productVideos []catalog.ProductVideo
	var videos []media.Video
	collection := cr.database.Collection(catalog.CollectionProductVideo)
	cursor, err := collection.Find(c, bson.M{"product_id": product.ID})
	if err != nil {
		return videos, err
	}

	err = cursor.All(c, &productVideos)
	if err != nil {
		return videos, err
	}

	var video media.Video
	collection = cr.database.Collection(media.CollectionVideo)
	for i := range productVideos {
		err = collection.FindOne(c, bson.M{"_id": productVideos[i].VideoID}).Decode(&video)
		if err == nil {
			videos = append(videos, video)
		}
	}

	return videos, err
}
