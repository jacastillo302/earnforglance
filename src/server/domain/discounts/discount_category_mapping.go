package domain

const (
	CollectionDiscountCategoryMapping = "discount_category_mappings"
)

// DiscountCategoryMapping represents a discount-category mapping class
type DiscountCategoryMapping struct {
	DiscountMapping
	EntityID int `bson:"entity_id"`
}
