package authflow

import (
	"strconv"

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

func (repo *UserRepository) InsertUser(userId string, phoneNo int64) *messages.AppError {
	//TODO wrong

	if err := repo.session.Query("INSERT INTO users (user_id, phone_no) VALUES (?, ?)", []string{userId, strconv.Itoa(int(phoneNo))}).Exec(); err != nil {
		return messages.InternalServerError("Unable to insert user")
	}
	return nil
}
