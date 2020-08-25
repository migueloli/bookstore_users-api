package mysqlutils

import (
	"fmt"
	"strings"

	"github.com/go-sql-driver/mysql"
	"github.com/migueloli/bookstore_users-api/utils/errors"
)

const (
	errorNoRows = "no rows in result set"
)

// ParseError process the error as a MySQL Error and convert to a errors.RestErr
func ParseError(err error) *errors.RestErr {
	sqlErr, ok := err.(*mysql.MySQLError)

	if ok {
		switch sqlErr.Number {
		case 1062:
			return errors.NewBadRequestError("Invalid data.")
		}
	}

	if strings.Contains(err.Error(), errorNoRows) {
		return errors.NewNotFoundError("No record matching given ID.")
	}

	return errors.NewInternalServerError(
		fmt.Sprintf("Error parsing database response."),
	)
}
