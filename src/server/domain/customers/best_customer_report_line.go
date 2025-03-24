package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionBestCustomerReportLine = "best_customer_report_lines"
)

// BestCustomerReportLine represents a best customer report line
type BestCustomerReportLine struct {
	ID         primitive.ObjectID `bson:"_id,omitempty"`
	CustomerID primitive.ObjectID `bson:"customer_id"`
	OrderTotal float64            `bson:"order_total"`
	OrderCount int                `bson:"order_count"`
}

type BestCustomerReportLineRepository interface {
	CreateMany(c context.Context, items []BestCustomerReportLine) error
	Create(c context.Context, best_customer_repor_line *BestCustomerReportLine) error
	Update(c context.Context, best_customer_repor_line *BestCustomerReportLine) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]BestCustomerReportLine, error)
	FetchByID(c context.Context, ID string) (BestCustomerReportLine, error)
}

type BestCustomerReportLineUsecase interface {
	CreateMany(c context.Context, items []BestCustomerReportLine) error
	FetchByID(c context.Context, ID string) (BestCustomerReportLine, error)
	Create(c context.Context, best_customer_repor_line *BestCustomerReportLine) error
	Update(c context.Context, best_customer_repor_line *BestCustomerReportLine) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]BestCustomerReportLine, error)
}
