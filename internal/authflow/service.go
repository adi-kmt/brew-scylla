package authflow

import (
	"github.com/adi-kmt/brew-scylla/internal/common/messages"
	"github.com/adi-kmt/brew-scylla/internal/domain"
	"github.com/google/uuid"
)

type authService struct {
	userPort domain.UserPort
}

func NewAuthService(userPort domain.UserPort) *authService {
	return &authService{
		userPort: userPort,
	}
}

func (s *authService) Login() (string, *messages.AppError) {
	return s.userPort.Login()
}

func (s *authService) Register(phoneNo int64) *messages.AppError {
	userId, err := uuid.NewV7()
	if err != nil {
		return messages.InternalServerError("Unable to generate user id")
	}
	return s.userPort.Register(userId.String(), phoneNo)
}

func (s *authService) GetUserDetailsByID(username string) (domain.UserDetailsDto, *messages.AppError) {
	return s.userPort.GetUserDetailsByID(username)
}
