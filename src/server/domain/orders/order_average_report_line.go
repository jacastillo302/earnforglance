package domain

// OrderAverageReportLine represents an order average report line
type OrderAverageReportLine struct {
	CountOrders               int     `bson:"count_orders"`
	SumShippingExclTax        float64 `bson:"sum_shipping_excl_tax"`
	OrderPaymentFeeExclTaxSum float64 `bson:"order_payment_fee_excl_tax_sum"`
	SumTax                    float64 `bson:"sum_tax"`
	SumOrders                 float64 `bson:"sum_orders"`
	SumRefundedAmount         float64 `bson:"sum_refunded_amount"`
}
