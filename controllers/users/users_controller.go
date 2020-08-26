package users

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/migueloli/bookstore_users-api/domain/users"
	"github.com/migueloli/bookstore_users-api/services"
	"github.com/migueloli/bookstore_users-api/utils/errors"
)

// CreateUser is the entry point for creating an user.
func CreateUser(c *gin.Context) {
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

// GetUser is the entry point for getting the user by id.
func GetUser(c *gin.Context) {
	userID, userErr := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if userErr != nil {
		err := errors.NewBadRequestError("User ID should be a number.")
		c.JSON(err.Status, err)
		return
	}

	user, getErr := services.GetUser(userID)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}

	c.JSON(http.StatusOK, user)
}

// UpdateUser is the entry point for updating the user by id.
func UpdateUser(c *gin.Context) {
	userID, userErr := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if userErr != nil {
		err := errors.NewBadRequestError("User ID should be a number.")
		c.JSON(err.Status, err)
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
