package domain

// BestsellersReportLine represents a best sellers report line
type BestsellersReportLine struct {
	ProductID     int     `bson:"product_id"`
	ProductName   string  `bson:"product_name"`
	TotalAmount   float64 `bson:"total_amount"`
	TotalQuantity int     `bson:"total_quantity"`
}
