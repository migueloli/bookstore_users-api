package services

import (
	"net/http"

	"github.com/migueloli/bookstore_users-api/domain/users"
	"github.com/migueloli/bookstore_users-api/utils/errors"
)

// CreateUser is a service to handle the user creation
func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	return &user, nil
	return nil, &errors.RestErr{
		Status: http.StatusInternalServerError,
	}
}
