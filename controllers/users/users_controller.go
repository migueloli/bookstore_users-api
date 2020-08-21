package users

import (
	"net/http"

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
	}
	c.JSON(http.StatusCreated, result)
}

// GetUser is the entry point for getting the user by id.
func GetUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "Implement me!")
}
