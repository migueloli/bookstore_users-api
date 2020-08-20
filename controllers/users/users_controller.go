package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateUser is the entry point for creating an user.
func CreateUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "Implement me!")
}

// GetUser is the entry point for getting the user by id.
func GetUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "Implement me!")
}
