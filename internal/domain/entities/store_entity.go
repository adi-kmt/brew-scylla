package entities

import "time"

type StoreEntity struct {
	City      string            `json:"city"`
	StoreName string            `json:"store_name"`
	Location  map[string]string `json:"location"`
	ClosesAt  time.Time         `json:"closes_at"`
}
