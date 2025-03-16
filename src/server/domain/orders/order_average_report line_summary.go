package domain

// OrderAverageReportLineSummary represents an order average report line summary
type OrderAverageReportLineSummary struct {
	OrderStatus          OrderStatus `bson:"order_status"`
	SumTodayOrders       float64     `bson:"sum_today_orders"`
	CountTodayOrders     int         `bson:"count_today_orders"`
	SumThisWeekOrders    float64     `bson:"sum_this_week_orders"`
	CountThisWeekOrders  int         `bson:"count_this_week_orders"`
	SumThisMonthOrders   float64     `bson:"sum_this_month_orders"`
	CountThisMonthOrders int         `bson:"count_this_month_orders"`
	SumThisYearOrders    float64     `bson:"sum_this_year_orders"`
	CountThisYearOrders  int         `bson:"count_this_year_orders"`
	SumAllTimeOrders     float64     `bson:"sum_all_time_orders"`
	CountAllTimeOrders   int         `bson:"count_all_time_orders"`
}
