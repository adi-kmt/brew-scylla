package authflow

import (
	"github.com/adi-kmt/brew-scylla/internal/common/messages"
	"github.com/adi-kmt/brew-scylla/internal/db"
	"github.com/adi-kmt/brew-scylla/internal/domain/entities"
	"github.com/scylladb/gocqlx/v3"
)

type UserRepository struct {
	session gocqlx.Session
}

func NewUserRepository(session gocqlx.Session) *UserRepository {
	return &UserRepository{
		session: session,
	}
}

func (repo *UserRepository) GetUserDetailsByID(userName string) (*entities.UserDetailsEntity, *messages.AppError) {
	var user entities.UserDetailsEntity
	err := db.GetUserDetailsByIDTable.SelectQuery(repo.session).Bind(userName).Select(&user)
	if err != nil {
		return nil, messages.InternalServerError("Unable to get user details")
	}
	return &user, nil
}

func (repo *UserRepository) InsertUser(userId string, phoneNo int64) (string, *messages.AppError) {
	userData := struct {
		UserID      string   `json:"user_id"`
		PhoneNo     int64    `json:"phone_no"`
		Coins       int      `json:"coins"`
		CouponsUsed []string `json:"coupons_used"`
	}{
		UserID:      userId,
		PhoneNo:     phoneNo,
		Coins:       50,
		CouponsUsed: []string{""},
	}

	err := db.GetUserDetailsByIDTable.InsertQuery(repo.session).BindStruct(userData).ExecRelease()
	if err != nil {
		return "", messages.InternalServerError("Unable to insert user")
	}
	return userId, nil
}

func (repo *UserRepository) UpdateUserDetails(userId string, userDetails *entities.UserDetailsEntity) *messages.AppError {
	return nil
}
