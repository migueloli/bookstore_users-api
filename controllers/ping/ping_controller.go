package ping

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Ping is a status check call.
func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
