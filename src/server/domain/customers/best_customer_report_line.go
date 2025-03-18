package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionBestCustomerReportLine = "best_customer_report_line"
)

// BestCustomerReportLine represents a best customer report line
type BestCustomerReportLine struct {
	ID         primitive.ObjectID `bson:"_id,omitempty"`
	CustomerID primitive.ObjectID `bson:"customer_id"`
	OrderTotal float64            `bson:"order_total"`
	OrderCount int                `bson:"order_count"`
}

type BestCustomerReportLineRepository interface {
	Create(c context.Context, best_customer_repor_line *BestCustomerReportLine) error
	Update(c context.Context, best_customer_repor_line *BestCustomerReportLine) error
	Delete(c context.Context, best_customer_repor_line *BestCustomerReportLine) error
	Fetch(c context.Context) ([]BestCustomerReportLine, error)
	FetchByID(c context.Context, best_customer_repor_lineID string) (BestCustomerReportLine, error)
}

type BestCustomerReportLineUsecase interface {
	FetchByID(c context.Context, best_customer_repor_lineID string) (BestCustomerReportLine, error)
	Create(c context.Context, best_customer_repor_line *BestCustomerReportLine) error
	Update(c context.Context, best_customer_repor_line *BestCustomerReportLine) error
	Delete(c context.Context, best_customer_repor_line *BestCustomerReportLine) error
	Fetch(c context.Context) ([]BestCustomerReportLine, error)
}
