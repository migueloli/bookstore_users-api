package app

import "github.com/gin-gonic/gin"

var (
	router = gin.Default()
)

// StartApplication configure and start the modules for de application.
func StartApplication() {
	mapUrls()
	router.Run(":8080")
}
