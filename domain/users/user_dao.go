package users

import (
	"errors"
	"fmt"
	"strings"

	"github.com/migueloli/bookstore_users-api/datasources/mysql/usersdb"
	"github.com/migueloli/bookstore_users-api/logger"
	"github.com/migueloli/bookstore_users-api/utils/mysqlutils"
	"github.com/migueloli/bookstore_utils-go/resterrors"
)

const (
	queryInsertUser              = "INSERT INTO users(first_name, last_name, email, date_created, status, password) VALUES (?, ?, ?, ?, ?, ?);"
	queryGetUser                 = "SELECT id, first_name, last_name, email, date_created, status FROM users WHERE id = ?;"
	queryUpdateUser              = "UPDATE users SET first_name = ?, last_name = ?, email = ? FROM users WHERE id = ?;"
	queryDeleteUser              = "DELETE FROM users WHERE id = ?;"
	queryFindUserByStatus        = "SELECT id, first_name, last_name, email, date_created, status FROM users WHERE status = ?;"
	queryFindUserByEmailPassword = "SELECT id, first_name, last_name, email, date_created, status FROM users WHERE email = ? AND password = ? AND status = ?;"
)

// Save the user in the database or return the RestErr.
func (user *User) Save() *resterrors.RestErr {
	stmt, err := usersdb.Client.Prepare(queryInsertUser)
	if err != nil {
		logger.Error("Error when trying to prepare the save user statement.", err)
		return resterrors.NewInternalServerError("Error when trying to prepare the save user statement.", errors.New("database error"))
	}

	defer stmt.Close()

	insertResult, saveErr := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated, user.Status, user.Password)
	if saveErr != nil {
		logger.Error("Error when trying to save user.", saveErr)
		return resterrors.NewInternalServerError("Error when trying to save user.", errors.New("database error"))
	}

	userID, err := insertResult.LastInsertId()
	if err != nil {
		logger.Error("Error when trying to get the last inserted userID.", err)
		return resterrors.NewInternalServerError("Error when trying to get the last inserted userID.", errors.New("database error"))
	}

	user.ID = userID

	return nil
}

// Get the user from the database or return a RestErr.
func (user *User) Get() *resterrors.RestErr {
	stmt, err := usersdb.Client.Prepare(queryGetUser)
	if err != nil {
		logger.Error("Error when trying to prepare the get user statement.", err)
		return resterrors.NewInternalServerError("Error when trying to prepare the get user statement.", errors.New("database error"))
	}

	defer stmt.Close()

	result := stmt.QueryRow(user.ID)
	if getErr := result.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status); getErr != nil {
		logger.Error("Error when trying to get user.", getErr)
		return resterrors.NewInternalServerError("Error when trying to get user.", errors.New("database error"))
	}

	return nil
}

// Update the user in the database or return the RestErr.
func (user *User) Update() *resterrors.RestErr {
	stmt, err := usersdb.Client.Prepare(queryUpdateUser)
	if err != nil {
		logger.Error("Error when trying to prepare the update user statement.", err)
		return resterrors.NewInternalServerError("Error when trying to prepare the update user statement.", errors.New("database error"))
	}

	defer stmt.Close()

	if _, err = stmt.Exec(user.FirstName, user.LastName, user.Email, user.ID); err != nil {
		logger.Error("Error when trying to update user.", err)
		return resterrors.NewInternalServerError("Error when trying to update user.", errors.New("database error"))
	}

	return nil
}

// Delete the user in the database or return the RestErr.
func (user *User) Delete() *resterrors.RestErr {
	stmt, err := usersdb.Client.Prepare(queryDeleteUser)
	if err != nil {
		logger.Error("Error when trying to prepare the delete user statement.", err)
		return resterrors.NewInternalServerError("Error when trying to prepare the delete user statement.", errors.New("database error"))
	}

	defer stmt.Close()

	if _, err = stmt.Exec(user.ID); err != nil {
		logger.Error("Error when trying to delete user.", err)
		return resterrors.NewInternalServerError("Error when trying to delete user.", errors.New("database error"))
	}

	return nil
}

// FindByStatus is a function to find the user using the status from the database or returning a RestErr
func (user *User) FindByStatus(status string) ([]User, *resterrors.RestErr) {
	stmt, err := usersdb.Client.Prepare(queryFindUserByStatus)
	if err != nil {
		logger.Error("Error when trying to prepare the find users by status statement.", err)
		return nil, resterrors.NewInternalServerError("Error when trying to prepare the find users by status statement.", errors.New("database error"))
	}

	defer stmt.Close()

	rows, err := stmt.Query(status)
	if err != nil {
		logger.Error("Error when trying to find users by status statement.", err)
		return nil, resterrors.NewInternalServerError("Error when trying to find users by status statement.", errors.New("database error"))
	}

	defer rows.Close()

	results := make([]User, 0)

	for rows.Next() {
		var result User
		if getErr := rows.Scan(&result.ID, &result.FirstName, &result.LastName, &result.Email, &result.DateCreated, &result.Status); getErr != nil {
			logger.Error("Error when trying to scan the user row into the user struct.", err)
			return nil, resterrors.NewInternalServerError("Error when trying to scan the user row into the user struct.", errors.New("database error"))
		}
		results = append(results, result)
	}

	if len(results) == 0 {
		return nil, resterrors.NewNotFoundError(fmt.Sprintf("No users matching status %s.", status))
	}

	return results, nil
}

// FindByEmailPassword the user from the database with a e-mail and password.
func (user *User) FindByEmailPassword() *resterrors.RestErr {
	stmt, err := usersdb.Client.Prepare(queryFindUserByEmailPassword)
	if err != nil {
		logger.Error("Error when trying to prepare the get user by e-mail and password statement", err)
		return resterrors.NewInternalServerError("Error when trying to prepare the get user by e-mail and password statement", errors.New("database error"))
	}

	defer stmt.Close()

	result := stmt.QueryRow(user.Email, user.Password, StatusActive)
	if getErr := result.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status); getErr != nil {
		if strings.Contains(getErr.Error(), mysqlutils.ErrorNoRows) {
			return resterrors.NewNotFoundError("Invalid user credentials.")
		}
		logger.Error("Error when trying to get user by e-mail and password.", getErr)
		return resterrors.NewInternalServerError("Error when trying to get user by e-mail and password.", errors.New("database error"))
	}

	return nil
}
