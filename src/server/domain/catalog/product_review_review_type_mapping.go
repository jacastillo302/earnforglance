package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionProductReviewReviewTypeMapping = "product_review_review_type_mappings"
)

// ProductReviewReviewTypeMapping represents a product review and review type mapping
type ProductReviewReviewTypeMapping struct {
	ID              primitive.ObjectID `bson:"_id,omitempty"`
	ProductReviewID primitive.ObjectID `bson:"product_review_id"`
	ReviewTypeID    primitive.ObjectID `bson:"review_type_id"`
	Rating          int                `bson:"rating"`
}

type ProductReviewReviewTypeMappingRepository interface {
	Create(c context.Context, product_review_review_type_mapping *ProductReviewReviewTypeMapping) error
	Update(c context.Context, product_review_review_type_mapping *ProductReviewReviewTypeMapping) error
	Delete(c context.Context, product_review_review_type_mapping *ProductReviewReviewTypeMapping) error
	Fetch(c context.Context) ([]ProductReviewReviewTypeMapping, error)
	FetchByID(c context.Context, ProductReviewReviewTypeMappingID string) (ProductReviewReviewTypeMapping, error)
}

type ProductReviewReviewTypeMappingUsecase interface {
	FetchByID(c context.Context, product_review_review_type_mappingID string) (ProductReviewReviewTypeMapping, error)
	Create(c context.Context, product_review_review_type_mapping *ProductReviewReviewTypeMapping) error
	Update(c context.Context, product_review_review_type_mapping *ProductReviewReviewTypeMapping) error
	Delete(c context.Context, product_review_review_type_mapping *ProductReviewReviewTypeMapping) error
	Fetch(c context.Context) ([]ProductReviewReviewTypeMapping, error)
}
