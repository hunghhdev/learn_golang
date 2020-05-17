package app

import (
	"users/src/logger"

	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

// StartApplication run app
func StartApplication() {
	mapUrls()

	logger.Info("about to start the application...")
	router.Run(":3000")
}
