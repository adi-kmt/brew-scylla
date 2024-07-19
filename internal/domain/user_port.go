package domain

import (
	"github.com/adi-kmt/brew-scylla/internal/common/messages"
	"github.com/adi-kmt/brew-scylla/internal/domain/entities"
)

type UserPort interface {
	Login() (string, *messages.AppError)
	Register(userName string, phoneNo int64) *messages.AppError
	GetUserDetailsByID(username string) (entities.UserDetailsEntity, *messages.AppError)
}
