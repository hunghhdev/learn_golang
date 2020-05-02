package users

import (
	"net/http"
	"users/src/domain/users"
	"users/src/services"
	"users/src/utils/errors"

	"github.com/gin-gonic/gin"
)

var (
	counter int
)

// CreateUser api
func CreateUser(c *gin.Context) {
	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
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

// GetUser api
func GetUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "impl me!")
}
