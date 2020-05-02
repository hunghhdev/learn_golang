package services

import (
	"net/http"
	"users/src/domain/users"
	"users/src/utils/errors"
)

// CreateUser just create
func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	return &user, nil
	return &user, &errors.RestErr{
		Status: http.StatusInternalServerError,
	}
}
