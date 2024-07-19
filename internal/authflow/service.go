package authflow

import (
	"github.com/adi-kmt/brew-scylla/internal/common/messages"
	"github.com/adi-kmt/brew-scylla/internal/domain"
	"github.com/adi-kmt/brew-scylla/internal/domain/entities"
	"github.com/google/uuid"
)

type AuthService struct {
	userPort domain.UserPort
}

func NewAuthService(userPort domain.UserPort) *AuthService {
	return &AuthService{
		userPort: userPort,
	}
}

func (s *AuthService) Login() (string, *messages.AppError) {
	return "", nil
}

func (s *AuthService) Register(phoneNo int64) *messages.AppError {
	userId, err := uuid.NewV7()
	if err != nil {
		return messages.InternalServerError("Unable to generate user id")
	}
	return s.userPort.InsertUser(userId.String(), phoneNo)
}

func (s *AuthService) GetUserDetailsByID(username string) (*entities.UserDetailsEntity, *messages.AppError) {
	return s.userPort.GetUserDetailsByID(username)
}
