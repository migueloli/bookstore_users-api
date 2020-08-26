package users

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/migueloli/bookstore_users-api/domain/users"
	"github.com/migueloli/bookstore_users-api/services"
	"github.com/migueloli/bookstore_users-api/utils/errors"
)

func getUserID(userIDParam string) (int64, *errors.RestErr) {
	userID, userErr := strconv.ParseInt(userIDParam, 10, 64)
	if userErr != nil {
		return 0, errors.NewBadRequestError("User ID should be a number.")
	}

	return userID, nil
}

// Create is the entry point for creating an user.
func Create(c *gin.Context) {
	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("Invalid JSON body.")
		c.JSON(restErr.Status, restErr)
		return
	}

	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}

	c.JSON(http.StatusCreated, result)
}

// Get is the entry point for getting the user by id.
func Get(c *gin.Context) {
	userID, idErr := getUserID(c.Param("user_id"))
	if idErr != nil {
		c.JSON(idErr.Status, idErr)
		return
	}

	user, getErr := services.GetUser(userID)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}

	c.JSON(http.StatusOK, user)
}

// Update is the entry point for updating the user by id.
func Update(c *gin.Context) {
	userID, idErr := getUserID(c.Param("user_id"))
	if idErr != nil {
		c.JSON(idErr.Status, idErr)
		return
	}

	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("Invalid JSON body.")
		c.JSON(restErr.Status, restErr)
		return
	}

	isPartial := c.Request.Method == http.MethodPatch

	user.ID = userID

	result, err := services.UpdateUser(isPartial, user)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusOK, result)
}

// Delete is the entry point for deleting the user by id.
func Delete(c *gin.Context) {
	userID, idErr := getUserID(c.Param("user_id"))
	if idErr != nil {
		c.JSON(idErr.Status, idErr)
		return
	}

	user := users.User{}

	user.ID = userID

	if err := services.DeleteUser(user); err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusOK, map[string]string{"status": "Deleted successfully."})
}
