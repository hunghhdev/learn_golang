package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	counter int
)

// CreateUser api
func CreateUser(c *gin.Context) {
	c.String(http.StatusNoContent, "impl me!")
}

// GetUser api
func GetUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "impl me!")
}
