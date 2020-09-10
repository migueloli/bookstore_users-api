package mysqlutils

import (
	"errors"
	"strings"

	"github.com/go-sql-driver/mysql"
	"github.com/migueloli/bookstore_utils-go/resterrors"
)

const (
	// ErrorNoRows is a message returned by the database to be used as comparission for identify the error.
	ErrorNoRows = "no rows in result set"
)

// ParseError process the error as a MySQL Error and convert to a errors.RestErr
func ParseError(err error) *resterrors.RestErr {
	sqlErr, ok := err.(*mysql.MySQLError)

	if ok {
		switch sqlErr.Number {
		case 1062:
			return resterrors.NewBadRequestError("Invalid data.")
		}
	}

	if strings.Contains(err.Error(), ErrorNoRows) {
		return resterrors.NewNotFoundError("No record matching given ID.")
	}

	return resterrors.NewInternalServerError(
		"Error parsing database response.",
		errors.New("database error"),
	)
}
