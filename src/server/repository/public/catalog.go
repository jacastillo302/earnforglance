package repository

import (
	"context"
	catalog "earnforglance/server/domain/catalog"
	media "earnforglance/server/domain/media"
	domain "earnforglance/server/domain/public"
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

func (cr *catalogRepository) GetProducts(c context.Context, filter domain.ProductRequest) ([]domain.ProductsResponse, error) {
	var result []domain.ProductsResponse
	var products []catalog.Product

	idHex, err := primitive.ObjectIDFromHex(filter.ID)
	if err == nil {
		var product catalog.Product

		collection := cr.database.Collection(catalog.CollectionProduct)
		err = collection.FindOne(c, bson.M{"_id": idHex, "deleted": false}).Decode(&product)
		if err != nil {
			return result, err
		}

		item, err := PrepareProduct(cr, c, product)
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

	for _, value := range filter.Filters {
		// "contains", "eq", etc.
		if value.Operator == "contains" {
			query[value.Field] = bson.M{"$regex": value.Value, "$options": "i"}
		} else {
			query[value.Field] = value.Value
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

	//fmt.Println("query", query)
	findOptions := options.Find().
		SetLimit(int64(filter.Limit)).
		SetSort(bson.D{{Key: "created_on_utc", Value: sortOrder}})

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
		item, err := PrepareProduct(cr, c, products[i])
		if err != nil {
			return result, err
		}
		items = append(items, item)
	}
	result = append(result, domain.ProductsResponse{Products: items})

	return result, err
}

func (cr *catalogRepository) GetProduct(c context.Context, ID string) (domain.ProductResponse, error) {

	var result domain.ProductResponse

	idHex, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return result, err
	}

	var product catalog.Product

	collection := cr.database.Collection(catalog.CollectionProduct)
	err = collection.FindOne(c, bson.M{"_id": idHex, "deleted": false}).Decode(&product)

	if err == nil {
		result, err = PrepareProduct(cr, c, product)
	}

	return result, err
}

func PrepareProduct(cr *catalogRepository, c context.Context, product catalog.Product) (domain.ProductResponse, error) {
	var result domain.ProductResponse

	result.Product = product

	var template catalog.ProductTemplate
	collection := cr.database.Collection(catalog.CollectionProductTemplate)
	err := collection.FindOne(c, bson.M{"_id": product.ProductTemplateID}).Decode(&template)
	if err == nil {
		result.Template = template
	}

	var productcategory []catalog.ProductCategory
	collection = cr.database.Collection(catalog.CollectionProductCategory)
	cursor, err := collection.Find(c, bson.M{"product_id": product.ID})
	if err != nil {
		return result, err
	}

	err = cursor.All(c, &productcategory)
	if err != nil {
		return result, err
	}

	var category catalog.Category
	collection = cr.database.Collection(catalog.CollectionCategory)
	for i := range productcategory {
		err = collection.FindOne(c, bson.M{"_id": productcategory[i].CategoryID}).Decode(&category)
		if err == nil {
			result.Categories = append(result.Categories, category)
		}
	}

	var productSpecificationAttribute []catalog.ProductSpecificationAttribute
	collection = cr.database.Collection(catalog.CollectionProductSpecificationAttribute)
	cursor, err = collection.Find(c, bson.M{"product_id": product.ID})
	if err != nil {
		return result, err
	}

	err = cursor.All(c, &productSpecificationAttribute)
	if err != nil {
		return result, err
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

			result.Specifications = append(result.Specifications, domain.SpecificationAttribute{Attribute: specificationAttribute, Options: specificationAttributeOptions, Group: specificationAttributeGroup})

			specificationAttributeOptions = nil
			specificationAttributeOptions = append(specificationAttributeOptions, specificationAttributeOption)
			specificationAttribute = specificationAttributeTemp
			specificationAttributeGroup = specificationAttributeGroupTemp

			bNewAttribute = false
			bSaveAttribute = true
		}

	}

	if bSaveAttribute {
		result.Specifications = append(result.Specifications, domain.SpecificationAttribute{Attribute: specificationAttribute, Options: specificationAttributeOptions, Group: specificationAttributeGroup})
		specificationAttributeOptions = nil
	}

	var productAttributeMap []catalog.ProductAttributeMapping
	collection = cr.database.Collection(catalog.CollectionProductAttributeMapping)
	cursor, err = collection.Find(c, bson.M{"product_id": product.ID})
	if err != nil {
		return result, err
	}

	err = cursor.All(c, &productAttributeMap)
	if err != nil {
		return result, err
	}

	for i := range productAttributeMap {

		var productAttribute catalog.ProductAttribute
		collection = cr.database.Collection(catalog.CollectionProductAttribute)
		collection.FindOne(c, bson.M{"_id": productAttributeMap[i].ProductAttributeID}).Decode(&productAttribute)

		var productAttributeValues []catalog.ProductAttributeValue
		collection = cr.database.Collection(catalog.CollectionProductAttributeValue)
		cursor, err = collection.Find(c, bson.M{"product_attribute_mapping_id": productAttributeMap[i].ID})
		if err != nil {
			return result, err
		}

		err = cursor.All(c, &productAttributeValues)
		if err != nil {
			return result, err
		}

		var productAttributeValue []domain.ProductAttributeValue
		for i := range productAttributeValues {

			var valuespictures []catalog.ProductAttributeValuePicture
			collection = cr.database.Collection(catalog.CollectionProductAttributeValuePicture)
			cursor, err = collection.Find(c, bson.M{"product_attribute_value_id": productAttributeValues[i].ID})
			if err != nil {
				return result, err
			}

			err = cursor.All(c, &valuespictures)
			if err != nil {
				return result, err
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
			return result, err
		}

		err = cursor.All(c, &catalogAttributeCombinations)
		if err != nil {
			return result, err
		}

		for f := range catalogAttributeCombinations {

			var catalogAttributeCombinationPictures []catalog.ProductAttributeCombinationPicture
			collection = cr.database.Collection(catalog.CollectionProductAttributeCombinationPicture)
			cursor, err = collection.Find(c, bson.M{"product_id": catalogAttributeCombinations[f].ProductID})
			if err != nil {
				return result, err
			}

			err = cursor.All(c, &catalogAttributeCombinationPictures)
			if err != nil {
				return result, err
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

		result.Attributes = append(result.Attributes, domain.ProductAttribute{Attribute: productAttribute, Values: productAttributeValue, Combinations: productAttributeCombinations})

	}

	var warehouse shipping.Warehouse
	collection = cr.database.Collection(shipping.CollectionWarehouse)
	err = collection.FindOne(c, bson.M{"_id": product.WarehouseID}).Decode(&warehouse)
	if err == nil {
		result.Warehouse.Warehouse = warehouse
	}

	var inventory catalog.ProductWarehouseInventory
	collection = cr.database.Collection(catalog.CollectionProductWarehouseInventory)
	err = collection.FindOne(c, bson.M{"_id": product.WarehouseID}).Decode(&inventory)
	if err == nil {
		result.Warehouse.Inventory = inventory
	}

	var deliveryDate shipping.DeliveryDate
	collection = cr.database.Collection(shipping.CollectionDeliveryDate)
	err = collection.FindOne(c, bson.M{"_id": product.DeliveryDateID}).Decode(&deliveryDate)
	if err == nil {
		result.DeliveryDate = deliveryDate
	}

	var availabilityRange shipping.ProductAvailabilityRange
	collection = cr.database.Collection(shipping.CollectionProductAvailabilityRange)
	err = collection.FindOne(c, bson.M{"_id": product.ProductAvailabilityRangeID}).Decode(&availabilityRange)
	if err == nil {
		result.Range = availabilityRange
	}

	var taxes tax.TaxCategory
	collection = cr.database.Collection(tax.CollectionTaxCategory)
	err = collection.FindOne(c, bson.M{"_id": product.TaxCategoryID}).Decode(&taxes)
	if err == nil {
		result.Tax = taxes
	}

	var vendo vendor.Vendor
	collection = cr.database.Collection(vendor.CollectionVendor)
	err = collection.FindOne(c, bson.M{"_id": product.VendorID}).Decode(&vendo)
	if err == nil {
		result.Vendor = vendo
	}

	var download media.Download
	collection = cr.database.Collection(media.CollectionDownload)
	err = collection.FindOne(c, bson.M{"_id": product.DownloadID}).Decode(&download)
	if err == nil {
		result.Download = &download
	} else {
		result.Download = nil
	}

	var tier []catalog.TierPrice
	collection = cr.database.Collection(catalog.CollectionTierPrice)
	cursor, err = collection.Find(c, bson.M{"product_id": product.ID})
	if err != nil {
		return result, err
	}

	err = cursor.All(c, &tier)
	if err != nil {
		return result, err
	}

	result.TierPrice = tier

	var crossproduct []catalog.CrossSellProduct
	collection = cr.database.Collection(catalog.CollectionCrossSellProduct)
	cursor, err = collection.Find(c, bson.M{"product_id1": product.ID})
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
			result.Relates = append(result.Relates, productcross1)
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
			result.Relates = append(result.Relates, productcross2)
		}
	}

	var relateproduct []catalog.RelatedProduct
	collection = cr.database.Collection(catalog.CollectionRelatedProduct)
	cursor, err = collection.Find(c, bson.M{"product_id1": product.ID})
	if err != nil {
		return result, err
	}

	err = cursor.All(c, &relateproduct)
	if err != nil {
		return result, err
	}

	var productrel1 catalog.Product
	collection = cr.database.Collection(catalog.CollectionProduct)
	for i := range relateproduct {
		err = collection.FindOne(c, bson.M{"_id": relateproduct[i].ProductID1}).Decode(&productrel1)
		if err == nil {
			result.Relates = append(result.Relates, productrel1)
		}
	}

	var relateproduct2 []catalog.RelatedProduct
	collection = cr.database.Collection(catalog.CollectionRelatedProduct)
	cursor, err = collection.Find(c, bson.M{"product_id2": product.ID})
	if err != nil {
		return result, err
	}

	err = cursor.All(c, &relateproduct2)
	if err != nil {
		return result, err
	}

	var productrel2 catalog.Product
	collection = cr.database.Collection(catalog.CollectionProduct)
	for i := range relateproduct2 {
		err = collection.FindOne(c, bson.M{"_id": relateproduct2[i].ProductID1}).Decode(&productrel2)
		if err == nil {
			result.Relates = append(result.Relates, productrel2)
		}
	}

	var productagmap []catalog.ProductProductTagMapping
	collection = cr.database.Collection(catalog.CollectionProductProductTagMapping)
	cursor, err = collection.Find(c, bson.M{"product_id": product.ID})
	if err != nil {
		return result, err
	}

	err = cursor.All(c, &productagmap)
	if err != nil {
		return result, err
	}

	var producta catalog.ProductTag
	collection = cr.database.Collection(catalog.CollectionProductTag)
	for i := range productagmap {
		err = collection.FindOne(c, bson.M{"_id": productagmap[i].ProductTagID}).Decode(&producta)
		if err == nil {
			result.Tags = append(result.Tags, producta)
		}
	}

	var manufacturerMaping []catalog.ProductManufacturer
	collection = cr.database.Collection(catalog.CollectionProductManufacturer)
	cursor, err = collection.Find(c, bson.M{"product_id": product.ID})
	if err != nil {
		return result, err
	}

	err = cursor.All(c, &manufacturerMaping)
	if err != nil {
		return result, err
	}

	var manufacturer catalog.Manufacturer
	collection = cr.database.Collection(catalog.CollectionManufacturer)
	for i := range manufacturerMaping {
		err = collection.FindOne(c, bson.M{"_id": manufacturerMaping[i].ManufacturerID}).Decode(&manufacturer)
		if err == nil {
			result.Manufacturers = append(result.Manufacturers, manufacturer)
		}
	}

	var productPictures []catalog.ProductPicture
	collection = cr.database.Collection(catalog.CollectionProductPicture)
	cursor, err = collection.Find(c, bson.M{"product_id": product.ID})
	if err != nil {
		return result, err
	}

	err = cursor.All(c, &productPictures)
	if err != nil {
		return result, err
	}

	var picture media.Picture
	collection = cr.database.Collection(media.CollectionPicture)
	for i := range productPictures {
		err = collection.FindOne(c, bson.M{"_id": productPictures[i].PictureID}).Decode(&picture)
		if err == nil {
			result.Pictures = append(result.Pictures, picture)
		}
	}

	var productVideos []catalog.ProductVideo
	collection = cr.database.Collection(catalog.CollectionProductVideo)
	cursor, err = collection.Find(c, bson.M{"product_id": product.ID})
	if err != nil {
		return result, err
	}

	err = cursor.All(c, &productVideos)
	if err != nil {
		return result, err
	}

	var video media.Video
	collection = cr.database.Collection(media.CollectionVideo)
	for i := range productVideos {
		err = collection.FindOne(c, bson.M{"_id": productVideos[i].VideoID}).Decode(&video)
		if err == nil {
			result.Videos = append(result.Videos, video)
		}
	}

	return result, err
}
