package app

import (
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

// StartApplication run app
func StartApplication() {
	mapUrls()
	router.Run(":3000")
}
