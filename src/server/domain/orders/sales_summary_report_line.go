package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionSalesSummaryReportLine = "sales_summary_report_lines"
)

// SalesSummaryReportLine represents sales summary report line
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
