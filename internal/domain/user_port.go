package domain

import "github.com/adi-kmt/brew-scylla/internal/common/messages"

type UserPort interface {
	Login() (string, *messages.AppError)
	Register(userName string, phoneNo int64) *messages.AppError // Add username through uuid
	GetUserDetailsByID(username string) (UserDetailsDto, *messages.AppError)
}
