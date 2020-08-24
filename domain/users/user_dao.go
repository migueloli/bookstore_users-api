package users

import (
	"fmt"

	"github.com/migueloli/bookstore_users-api/utils/errors"
)

var (
	usersDB = make(map[int64]*User)
)

// Get the user from the database or return a RestErr.
func (user *User) Get() *errors.RestErr {
	result := usersDB[user.ID]
	if result == nil {
		return errors.NewNotFoundError(
			fmt.Sprintf("User %d not found.", user.ID),
		)
	}

	user.ID = result.ID
	user.FirstName = result.FirstName
	user.LastName = result.LastName
	user.Email = result.Email
	user.DateCreated = result.DateCreated

	return nil
}

// Save the user in the database or return the RestErr.
func (user *User) Save() *errors.RestErr {
	current := usersDB[user.ID]
	if current != nil {
		if user.Email == current.Email {
			return errors.NewBadRequestError(
				fmt.Sprintf("Email %s already exists", user.Email),
			)
		}
		return errors.NewBadRequestError(
			fmt.Sprintf("User %d already exists", user.ID),
		)
	}

	usersDB[user.ID] = user

	return nil
}
