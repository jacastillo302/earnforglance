package domain

// OrderByCountryReportLine represents an "order by country" report line
type OrderByCountryReportLine struct {
	CountryID   *int    `bson:"country_id,omitempty"`
	TotalOrders int     `bson:"total_orders"`
	SumOrders   float64 `bson:"sum_orders"`
}
