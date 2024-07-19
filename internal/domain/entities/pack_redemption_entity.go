package entities

import "time"

type PackRedemptionEntity struct {
	Username            string    `json:"username"`
	StoreName           string    `json:"store_name"`
	OrderID             string    `json:"order_id"`
	PackName            string    `json:"pack_name"`
	OrderItemsRemaining int       `json:"order_items_remaining"`
	ExpiryTimestamp     time.Time `json:"expiry_timestamp"`
}
