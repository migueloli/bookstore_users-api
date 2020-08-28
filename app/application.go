package app

import (
	"github.com/gin-gonic/gin"
	"github.com/migueloli/bookstore_users-api/logger"
)

var (
	router = gin.Default()
)

// StartApplication configure and start the modules for de application.
func StartApplication() {
	mapUrls()

	logger.Info("Starting application...")
	router.Run(":8080")
}
