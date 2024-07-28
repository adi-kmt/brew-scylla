package authflow

import (
	"github.com/adi-kmt/brew-scylla/internal/common/messages"
	"github.com/adi-kmt/brew-scylla/internal/db"
	"github.com/adi-kmt/brew-scylla/internal/domain/entities"
	"github.com/gocql/gocql"
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
	var user UserDetailsEntity
	err := db.GetUserDetailsByIDTable.SelectQuery(repo.session).Bind(userName).Select(&user)
	if err != nil {
		return nil, messages.InternalServerError("Unable to get user details")
	}
	return &entities.UserDetailsEntity{
		UserId:      user.UserID.String(),
		PhoneNo:     user.PhoneNo,
		Coins:       user.Coins,
		CouponsUsed: user.CouponsUsed,
	}, nil
}

func (repo *UserRepository) InsertUser(userId string, phoneNo int64) (string, *messages.AppError) {
	userUUID, err0 := gocql.ParseUUID(userId)
	if err0 != nil {
		return "", messages.BadRequest("Invalid user id")
	}

	userData := UserDetailsEntity{
		UserID:      userUUID,
		PhoneNo:     phoneNo,
		Coins:       50,
		CouponsUsed: []string{},
	}

	err := db.GetUserDetailsByIDTable.InsertQuery(repo.session).BindStruct(userData).ExecRelease()
	if err != nil {
		return "", messages.InternalServerError("Unable to insert user")
	}
	return userId, nil
}

func (repo *UserRepository) UpdateUserDetails(userId string, userDetails *entities.UserDetailsEntity) *messages.AppError {
	userUUID, err0 := gocql.ParseUUID(userId)
	if err0 != nil {
		return messages.BadRequest("Invalid user id")
	}

	err := db.GetUserDetailsByIDTable.UpdateQuery(repo.session).BindStruct(userDetails).Bind(userUUID).ExecRelease()
	if err != nil {
		return messages.InternalServerError("Unable to update user details")
	}
	return nil
}
