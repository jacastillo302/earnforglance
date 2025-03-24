package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionBestSellersReportLine = "best_sellers_report_line"
)

// BestSellersReportLine represents a best sellers report line
type BestSellersReportLine struct {
	ProductID     primitive.ObjectID `bson:"product_id"`
	ProductName   string             `bson:"product_name"`
	TotalAmount   float64            `bson:"total_amount"`
	TotalQuantity int                `bson:"total_quantity"`
}

type BestSellersReportLineRepository interface {
	CreateMany(c context.Context, items []BestSellersReportLine) error
	Create(c context.Context, best_sellers_report_line *BestSellersReportLine) error
	Update(c context.Context, best_sellers_report_line *BestSellersReportLine) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]BestSellersReportLine, error)
	FetchByID(c context.Context, ID string) (BestSellersReportLine, error)
}

type BestSellersReportLineUsecase interface {
	FetchByID(c context.Context, ID string) (BestSellersReportLine, error)
	Create(c context.Context, best_sellers_report_line *BestSellersReportLine) error
	Update(c context.Context, best_sellers_report_line *BestSellersReportLine) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]BestSellersReportLine, error)
}
