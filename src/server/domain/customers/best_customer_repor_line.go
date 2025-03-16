package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

const (
	CollectionBestCustomerReportLine = "best_customer_report_line"
)

// BestCustomerReportLine represents a best customer report line
type BestCustomerReportLine struct {
	ID         primitive.ObjectID `bson:"_id,omitempty"`
	CustomerID int                `bson:"customer_id"`
	OrderTotal float64            `bson:"order_total"`
	OrderCount int                `bson:"order_count"`
}
