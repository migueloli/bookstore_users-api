package users

import (
	"fmt"

	"github.com/migueloli/bookstore_users-api/datasources/mysql/usersdb"
	"github.com/migueloli/bookstore_users-api/logger"
	"github.com/migueloli/bookstore_users-api/utils/errors"
)

const (
	queryInsertUser       = "INSERT INTO users(first_name, last_name, email, date_created, status, password) VALUES (?, ?, ?, ?, ?, ?);"
	queryGetUser          = "SELECT id, first_name, last_name, email, date_created, status FROM users WHERE id = ?;"
	queryUpdateUser       = "UPDATE users SET first_name = ?, last_name = ?, email = ? FROM users WHERE id = ?;"
	queryDeleteUser       = "DELETE FROM users WHERE id = ?;"
	queryFindUserByStatus = "SELECT id, first_name, last_name, email, date_created, status FROM users WHERE status = ?;"
)

// Save the user in the database or return the RestErr.
func (user *User) Save() *errors.RestErr {
	stmt, err := usersdb.Client.Prepare(queryInsertUser)
	if err != nil {
		logger.Error("Error when trying to prepare the save user statement", err)
		return errors.NewInternalServerError("Database error")
	}

	defer stmt.Close()

	insertResult, saveErr := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated, user.Status, user.Password)
	if saveErr != nil {
		logger.Error("Error when trying to save user", saveErr)
		return errors.NewInternalServerError("Database error")
	}

	userID, err := insertResult.LastInsertId()
	if err != nil {
		logger.Error("Error when trying to get the last inserted userID", err)
		return errors.NewInternalServerError("Database error")
	}

	user.ID = userID

	return nil
}

// Get the user from the database or return a RestErr.
func (user *User) Get() *errors.RestErr {
	stmt, err := usersdb.Client.Prepare(queryGetUser)
	if err != nil {
		logger.Error("Error when trying to prepare the get user statement", err)
		return errors.NewInternalServerError("Database error")
	}

	defer stmt.Close()

	result := stmt.QueryRow(user.ID)
	if getErr := result.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status); getErr != nil {
		logger.Error("Error when trying to get user", getErr)
		return errors.NewInternalServerError("Database error")
	}

	return nil
}

// Update the user in the database or return the RestErr.
func (user *User) Update() *errors.RestErr {
	stmt, err := usersdb.Client.Prepare(queryUpdateUser)
	if err != nil {
		logger.Error("Error when trying to prepare the update user statement", err)
		return errors.NewInternalServerError("Database error")
	}

	defer stmt.Close()

	if _, err = stmt.Exec(user.FirstName, user.LastName, user.Email, user.ID); err != nil {
		logger.Error("Error when trying to update user", err)
		return errors.NewInternalServerError("Database error")
	}

	return nil
}

// Delete the user in the database or return the RestErr.
func (user *User) Delete() *errors.RestErr {
	stmt, err := usersdb.Client.Prepare(queryDeleteUser)
	if err != nil {
		logger.Error("Error when trying to prepare the delete user statement", err)
		return errors.NewInternalServerError("Database error")
	}

	defer stmt.Close()

	if _, err = stmt.Exec(user.ID); err != nil {
		logger.Error("Error when trying to delete user", err)
		return errors.NewInternalServerError("Database error")
	}

	return nil
}

// FindByStatus is a function to find the user using the status from the database or returning a RestErr
func (user *User) FindByStatus(status string) ([]User, *errors.RestErr) {
	stmt, err := usersdb.Client.Prepare(queryFindUserByStatus)
	if err != nil {
		logger.Error("Error when trying to prepare the find users by status statement", err)
		return nil, errors.NewInternalServerError("Database error")
	}

	defer stmt.Close()

	rows, err := stmt.Query(status)
	if err != nil {
		logger.Error("Error when trying to find users by status statement", err)
		return nil, errors.NewInternalServerError("Database error")
	}

	defer rows.Close()

	results := make([]User, 0)

	for rows.Next() {
		var result User
		if getErr := rows.Scan(&result.ID, &result.FirstName, &result.LastName, &result.Email, &result.DateCreated, &result.Status); getErr != nil {
			logger.Error("Error when trying to scan the user row into the user struct", err)
			return nil, errors.NewInternalServerError("Database error")
		}
		results = append(results, result)
	}

	if len(results) == 0 {
		return nil, errors.NewNotFoundError(fmt.Sprintf("No users matching status %s.", status))
	}

	return results, nil
}
