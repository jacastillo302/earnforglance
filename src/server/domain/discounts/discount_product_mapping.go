package domain

const (
	CollectionDiscountProductMapping = "discount_product_mappings"
)

// DiscountProductMapping represents a discount-product mapping class
type DiscountProductMapping struct {
	DiscountMapping
	EntityID int `bson:"entity_id"`
}
