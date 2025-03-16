package domain

const (
	CollectionDiscountManufacturerMapping = "discount_manufacturer_mappings"
)

// DiscountManufacturerMapping represents a discount-manufacturer mapping class
type DiscountManufacturerMapping struct {
	DiscountMapping
	EntityID int `bson:"entity_id"`
}
