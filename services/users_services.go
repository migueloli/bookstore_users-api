package services

import (
	"github.com/migueloli/bookstore_users-api/domain/users"
	"github.com/migueloli/bookstore_users-api/utils/cryptoutils"
	"github.com/migueloli/bookstore_users-api/utils/dateutils"
	"github.com/migueloli/bookstore_users-api/utils/errors"
)

var (
	// UsersService is the access point to the usersServiceInterface as usersService struct.
	UsersService usersServiceInterface = &usersService{}
)

type usersService struct{}

type usersServiceInterface interface {
	CreateUser(users.User) (*users.User, *errors.RestErr)
	GetUser(int64) (*users.User, *errors.RestErr)
	UpdateUser(bool, users.User) (*users.User, *errors.RestErr)
	DeleteUser(int64) *errors.RestErr
	SearchUser(string) (users.Users, *errors.RestErr)
	LoginUser(users.UserLoginRequest) (*users.User, *errors.RestErr)
}

// CreateUser is a service to handle the user creation
func (s *usersService) CreateUser(user users.User) (*users.User, *errors.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}

	user.Status = users.StatusActive
	user.DateCreated = dateutils.GetNowDBString()
	user.Password = cryptoutils.GetMd5(user.Password)
	if err := user.Save(); err != nil {
		return nil, err
	}

	return &user, nil
}

// GetUser is a service to handle the user recover
func (s *usersService) GetUser(userID int64) (*users.User, *errors.RestErr) {
	if userID <= 0 {
		return nil, errors.NewBadRequestError("User ID has to be greater than 0.")
	}

	result := &users.User{ID: userID}
	if err := result.Get(); err != nil {
		return nil, err
	}

	return result, nil
}

// UpdateUser is a service to handle the user updating
func (s *usersService) UpdateUser(isPartial bool, user users.User) (*users.User, *errors.RestErr) {
	current, err := s.GetUser(user.ID)
	if err != nil {
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

	if err := current.Validate(); err != nil {
		return nil, err
	}

	if err := current.Update(); err != nil {
		return nil, err
	}

	return &user, nil
}

// DeleteUser is a service to handle the user recover
func (s *usersService) DeleteUser(userID int64) *errors.RestErr {
	if userID <= 0 {
		return errors.NewBadRequestError("User ID has to be greater than 0.")
	}

	user := &users.User{ID: userID}
	return user.Delete()
}

// SearchUser is a service to handle the user recover using params
func (s *usersService) SearchUser(status string) (users.Users, *errors.RestErr) {
	dao := &users.User{}
	return dao.FindByStatus(status)
}

// LoginUser is a service to handle the user login
func (s *usersService) LoginUser(request users.UserLoginRequest) (*users.User, *errors.RestErr) {
	dao := &users.User{
		Email:    request.Email,
		Password: cryptoutils.GetMd5(request.Password),
	}
	if err := dao.FindByEmailPassword(); err != nil {
		return nil, err
	}
	return dao, nil
}
