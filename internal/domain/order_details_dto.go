package domain

import "time"

type OrderDetailsDTO struct {
	Username           string      `json:"username"`
	OrderID            string      `json:"order_id"`
	ProductName        string      `json:"product_name"`
	ProductQuantity    int         `json:"product_quantity"`
	ProductPrice       float64     `json:"product_price"`
	OrderStatus        OrderStatus `json:"order_status"`
	OrderTimestamp     time.Time   `json:"order_timestamp"`
	OrderSubTotal      float64     `json:"order_sub_total"`
	DiscountPercentage float64     `json:"discount_percentage"`
	OrderTotal         float64     `json:"order_total"`
	PackName           string      `json:"pack_name"`
	IsPack             bool        `json:"is_pack"`
}
