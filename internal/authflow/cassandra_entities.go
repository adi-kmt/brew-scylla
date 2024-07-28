package authflow

import "github.com/gocql/gocql"

type UserDetailsEntity struct {
	UserID      gocql.UUID `json:"user_id"`
	PhoneNo     int64      `json:"phone_no"`
	Coins       int        `json:"coins"`
	CouponsUsed []string   `json:"coupons_used"`
}
