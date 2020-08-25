package services

import (
	"github.com/migueloli/bookstore_users-api/domain/users"
	"github.com/migueloli/bookstore_users-api/utils/errors"
)

// CreateUser is a service to handle the user creation
func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}

	if err := user.Save(); err != nil {
		return nil, err
	}

	return &user, nil
}

// GetUser is a service to handle the user recover
func GetUser(userID int64) (*users.User, *errors.RestErr) {
	if userID <= 0 {
		return nil, errors.NewBadRequestError("User ID has to be greater than 0.")
	}

	result := &users.User{ID: userID}
	if err := result.Get(); err != nil {
		return nil, err
	}

	return result, nil
}
