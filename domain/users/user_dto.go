package users

import (
	"strings"

	"github.com/migueloli/bookstore_users-api/utils/errors"
)

// User is the base of this domain
type User struct {
	ID          int64  `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	DateCreated string `json:"date_created"`
}

// Validate is used to verify if the user struct has the obligated fields
// are correctly fulfilled
func (user *User) Validate() *errors.RestErr {
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	if user.Email == "" {
		return errors.NewBadRequestError("Invalid e-mail address.")
	}
	return nil
}
