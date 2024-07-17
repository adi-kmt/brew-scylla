package domain

import "time"

type CouponCodeDTO struct {
	StoreName  string    `json:"store_name"`
	CuponCode  string    `json:"cupon_code"`
	ExpiryDate time.Time `json:"expiry_date"`
	CuponImage string    `json:"cupon_image"`
}
