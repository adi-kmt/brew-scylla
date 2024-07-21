package entities

type UserDetailsEntity struct {
	UserId      string   `json:"user_id"`
	PhoneNo     int64    `json:"phone_no"`
	Coins       int      `json:"coins"`
	CouponsUsed []string `json:"coupons"`
}
