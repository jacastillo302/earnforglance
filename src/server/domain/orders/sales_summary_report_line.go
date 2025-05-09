package domain

import (
	"context" // added context library
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

const (
	CollectionSalesSummaryReportLine = "sales_summary_report_lines"
)

// SalesSummaryReportLine represents sales summary report line.
type SalesSummaryReportLine struct {
	ID             bson.ObjectID `bson:"_id,omitempty"`
	Summary        string        `bson:"summary"`
	SummaryDate    time.Time     `bson:"summary_date"`
	NumberOfOrders int           `bson:"number_of_orders"`
	Profit         float64       `bson:"profit"`
	ProfitStr      string        `bson:"profit_str"`
	Shipping       string        `bson:"shipping"`
	Tax            string        `bson:"tax"`
	OrderTotal     string        `bson:"order_total"`
	SummaryType    int           `bson:"summary_type"`
}

// SalesSummaryReportLineRepository represents the repository interface for SalesSummaryReportLine
type SalesSummaryReportLineRepository interface {
	CreateMany(c context.Context, items []SalesSummaryReportLine) error
	Create(c context.Context, sales_summary_report_line *SalesSummaryReportLine) error
	Update(c context.Context, sales_summary_report_line *SalesSummaryReportLine) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]SalesSummaryReportLine, error)
	FetchByID(c context.Context, ID string) (SalesSummaryReportLine, error)
}

// SalesSummaryReportLineUsecase represents the use case interface for SalesSummaryReportLine
type SalesSummaryReportLineUsecase interface {
	CreateMany(c context.Context, items []SalesSummaryReportLine) error
	FetchByID(c context.Context, ID string) (SalesSummaryReportLine, error)
	Create(c context.Context, sales_summary_report_line *SalesSummaryReportLine) error
	Update(c context.Context, sales_summary_report_line *SalesSummaryReportLine) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]SalesSummaryReportLine, error)
}
