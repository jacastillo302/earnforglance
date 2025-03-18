package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// BestsellersReportLine represents a best sellers report line
type BestsellersReportLine struct {
	ProductID     primitive.ObjectID `bson:"product_id"`
	ProductName   string             `bson:"product_name"`
	TotalAmount   float64            `bson:"total_amount"`
	TotalQuantity int                `bson:"total_quantity"`
}

type BestsellersReportLineRepository interface {
	Create(c context.Context, best_sellers_report_line *BestsellersReportLine) error
	Update(c context.Context, best_sellers_report_line *BestsellersReportLine) error
	Delete(c context.Context, best_sellers_report_line *BestsellersReportLine) error
	Fetch(c context.Context) ([]BestsellersReportLine, error)
	FetchByID(c context.Context, best_sellers_report_lineID string) (BestsellersReportLine, error)
}

type BestsellersReportLineUsecase interface {
	FetchByID(c context.Context, best_sellers_report_lineID string) (BestsellersReportLine, error)
	Create(c context.Context, best_sellers_report_line *BestsellersReportLine) error
	Update(c context.Context, best_sellers_report_line *BestsellersReportLine) error
	Delete(c context.Context, best_sellers_report_line *BestsellersReportLine) error
	Fetch(c context.Context) ([]BestsellersReportLine, error)
}
