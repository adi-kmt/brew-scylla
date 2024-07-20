package domain

import (
	"github.com/adi-kmt/brew-scylla/internal/common/messages"
	"github.com/adi-kmt/brew-scylla/internal/domain/entities"
)

type UserPort interface {
	InsertUser(userName string, phoneNo int64) (string, *messages.AppError)
	GetUserDetailsByID(username string) (*entities.UserDetailsEntity, *messages.AppError)
	UpdateUserDetails(userId string, userDetails *entities.UserDetailsEntity) *messages.AppError
}
