package domain

import "github.com/adi-kmt/brew-scylla/internal/common/messages"

type UserPort interface {
	// Use magic link to login/register
	Register(phoneNo int64) *messages.AppError // Add username through uuid
	GetUserDetailsByID(username string) (UserDetailsDto, *messages.AppError)
}
