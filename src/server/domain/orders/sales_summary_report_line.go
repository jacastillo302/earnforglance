package domain

import (
	"context" // added context library
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionSalesSummaryReportLine = "sales_summary_report_lines"
)

// SalesSummaryReportLine represents sales summary report line.
type SalesSummaryReportLine struct {
	ID             primitive.ObjectID `bson:"_id,omitempty"`
	Summary        string             `bson:"summary"`
	SummaryDate    time.Time          `bson:"summary_date"`
	NumberOfOrders int                `bson:"number_of_orders"`
	Profit         float64            `bson:"profit"`
	ProfitStr      string             `bson:"profit_str"`
	Shipping       string             `bson:"shipping"`
	Tax            string             `bson:"tax"`
	OrderTotal     string             `bson:"order_total"`
	SummaryType    int                `bson:"summary_type"`
}

// SalesSummaryReportLineRepository represents the repository interface for SalesSummaryReportLine
type SalesSummaryReportLineRepository interface {
	Create(c context.Context, sales_summary_report_line *SalesSummaryReportLine) error
	Update(c context.Context, sales_summary_report_line *SalesSummaryReportLine) error
	Delete(c context.Context, sales_summary_report_line *SalesSummaryReportLine) error
	Fetch(c context.Context) ([]SalesSummaryReportLine, error)
	FetchByID(c context.Context, sales_summary_report_lineID string) (SalesSummaryReportLine, error)
}

// SalesSummaryReportLineUsecase represents the use case interface for SalesSummaryReportLine
type SalesSummaryReportLineUsecase interface {
	FetchByID(c context.Context, sales_summary_report_lineID string) (SalesSummaryReportLine, error)
	Create(c context.Context, sales_summary_report_line *SalesSummaryReportLine) error
	Update(c context.Context, sales_summary_report_line *SalesSummaryReportLine) error
	Delete(c context.Context, sales_summary_report_line *SalesSummaryReportLine) error
	Fetch(c context.Context) ([]SalesSummaryReportLine, error)
}
