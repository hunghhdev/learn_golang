package services

import (
	"users/src/domain/users"
	"users/src/utils/crypto_utils"
	"users/src/utils/date_utils"
	"users/src/utils/errors"
)

var (
	UsersService usersServiceInterface = &usersService{}
)

type usersService struct{}

type usersServiceInterface interface {
	GetUser(int64) (*users.User, *errors.RestErr)
	CreateUser(users.User) (*users.User, *errors.RestErr)
	UpdateUser(bool, users.User) (*users.User, *errors.RestErr)
	DeleteUser(int64) *errors.RestErr
	SearchUser(string) (users.Users, *errors.RestErr)
	LoginUser(users.LoginRequest) (*users.User, *errors.RestErr)
}

// GetUser just create
func (s *usersService) GetUser(userID int64) (*users.User, *errors.RestErr) {
	result := &users.User{ID: userID}
	if err := result.Get(); err != nil {
		return nil, err
	}
	return result, nil
}

// CreateUser just create
func (s *usersService) CreateUser(user users.User) (*users.User, *errors.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}

	user.Status = users.StatusActive
	user.DateCreated = date_utils.GetNowDBFormat()
	user.Password = crypto_utils.GetMd5(user.Password)
	if err := user.Save(); err != nil {
		return nil, err
	}
	return &user, nil
}

// UpdateUser just update
func (s *usersService) UpdateUser(isPartial bool, user users.User) (*users.User, *errors.RestErr) {
	current, err := s.GetUser(user.ID)
	if err != nil {
		return nil, err
	}

	if err := user.Validate(); err != nil {
		return nil, err
	}

	if isPartial {
		if user.FirstName != "" {
			current.FirstName = user.FirstName
		}
		if user.LastName != "" {
			current.LastName = user.LastName
		}
		if user.Email != "" {
			current.Email = user.Email
		}
	} else {
		current.FirstName = user.FirstName
		current.LastName = user.LastName
		current.Email = user.Email
	}

	if err := current.Update(); err != nil {
		return nil, err
	}
	return current, nil
}

// DeleteUser just delete
func (s *usersService) DeleteUser(userID int64) *errors.RestErr {
	user := &users.User{ID: userID}
	return user.Delete()
}

// Search just find by status
func (s *usersService) SearchUser(status string) (users.Users, *errors.RestErr) {
	dao := &users.User{}
	return dao.FindByStatus(status)
}

// Login
func (s *usersService) LoginUser(request users.LoginRequest) (*users.User, *errors.RestErr) {
	dao := &users.User{
		Email:    request.Email,
		Password: crypto_utils.GetMd5(request.Password),
	}
	if err := dao.FindByEmailAndPassword(); err != nil {
		return nil, err
	}
	return dao, nil
}
