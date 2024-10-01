package models

type MetricsResponse struct {
	CarrierName          string  `json:"carrier_name"`
	TotalResults         int     `json:"total_results"`
	TotalShippingPrice   float64 `json:"total_shipping_price"`
	AveragePriceShipping float64 `json:"average_price_shipping"`
}

type OverallMetrics struct {
	TotalCarriers         int               `json:"total_carriers"`
	CheapestShippingPrice float64           `json:"cheapest_shipping_price"`
	ShippingMostExpensive float64           `json:"shipping_most_expensive"`
	MetricsByCarrier      []MetricsResponse `json:"metrics_by_carrier"`
}
