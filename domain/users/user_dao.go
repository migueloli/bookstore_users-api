package users

import (
	"github.com/migueloli/bookstore_users-api/datasources/mysql/usersdb"
	"github.com/migueloli/bookstore_users-api/utils/dateutils"
	"github.com/migueloli/bookstore_users-api/utils/errors"
	"github.com/migueloli/bookstore_users-api/utils/mysqlutils"
)

const (
	queryInsertUser = "INSERT INTO users(first_name, last_name, email, date_created) VALUES (?, ?, ?, ?);"
	queryGetUser    = "SELECT id, first_name, last_name, email, date_created FROM users WHERE id = ?;"
	queryUpdateUser = "UPDATE user SET first_name = ?, last_name = ?, email = ? FROM users WHERE id = ?;"
	queryDeleteUser = "DELETE FROM user WHERE id = ?;"
)

// Save the user in the database or return the RestErr.
func (user *User) Save() *errors.RestErr {
	stmt, err := usersdb.Client.Prepare(queryInsertUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}

	defer stmt.Close()

	user.DateCreated = dateutils.GetNowString()

	insertResult, saveErr := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated)
	if saveErr != nil {
		return mysqlutils.ParseError(saveErr)
	}

	userID, err := insertResult.LastInsertId()
	if err != nil {
		mysqlutils.ParseError(err)
	}

	user.ID = userID

	return nil
}

// Get the user from the database or return a RestErr.
func (user *User) Get() *errors.RestErr {
	stmt, err := usersdb.Client.Prepare(queryGetUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}

	defer stmt.Close()

	result := stmt.QueryRow(user.ID)
	if getErr := result.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated); getErr != nil {
		mysqlutils.ParseError(getErr)
	}

	return nil
}

// Update the user in the database or return the RestErr.
func (user *User) Update() *errors.RestErr {
	stmt, err := usersdb.Client.Prepare(queryUpdateUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}

	defer stmt.Close()

	if _, err = stmt.Exec(user.FirstName, user.LastName, user.Email, user.ID); err != nil {
		return mysqlutils.ParseError(err)
	}

	return nil
}

// Delete the user in the database or return the RestErr.
func (user *User) Delete() *errors.RestErr {
	stmt, err := usersdb.Client.Prepare(queryDeleteUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}

	defer stmt.Close()

	if _, err = stmt.Exec(user.ID); err != nil {
		return mysqlutils.ParseError(err)
	}

	return nil
}
