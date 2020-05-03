package users

import (
	"fmt"
	users_db "users/src/datasources/mysql/users"
	"users/src/utils/date_utils"
	"users/src/utils/errors"
)

var (
	usersDB = make(map[int64]*User)
)

// Get user by id
func (user *User) Get() *errors.RestErr {
	if err := users_db.Client.Ping(); err != nil {
		panic(err)
	}
	result := usersDB[user.ID]
	if result == nil {
		return errors.NewNotFoundError(fmt.Sprintf("user %d not found", user.ID))
	}

	user.ID = result.ID
	user.FirstName = result.FirstName
	user.LastName = result.LastName
	user.Email = result.Email
	user.DateCreated = result.DateCreated

	return nil
}

// Save user
func (user *User) Save() *errors.RestErr {
	current := usersDB[user.ID]
	if current != nil {
		if current.Email == user.Email {
			return errors.NewBadRequestError(fmt.Sprintf("Email %s already registered", user.Email))
		}
		return errors.NewBadRequestError(fmt.Sprintf("user %d already exists", user.ID))
	}
	user.DateCreated = date_utils.GetNowString()

	usersDB[user.ID] = user
	return nil
}
