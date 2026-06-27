package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (uc *userControllerInterface) DeleteUser(c *gin.Context) {
	userID := c.Param("userId")
	if err := uc.service.DeleteUser(userID); err != nil {
		c.JSON(err.Code, err)
		return
	}

	c.Status(http.StatusNoContent)
}
