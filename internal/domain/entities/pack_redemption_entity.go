package entities

import "time"

type PackRedemptionEntity struct {
	UserId              string    `json:"user_id"`
	StoreName           string    `json:"store_name"`
	OrderID             string    `json:"order_id"`
	PackName            string    `json:"pack_name"`
	OrderItemsRemaining int       `json:"order_items_remaining"`
	ExpiryTimestamp     time.Time `json:"expiry_timestamp"`
}
