package domain

type UserDetailsDto struct {
	Username string `json:"username"`
	PhoneNo  int64  `json:"phone_no"`
	Coins    int    `json:"coins"`
}
