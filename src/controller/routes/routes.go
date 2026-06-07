package routes

import (
	"github.com/CauaMarques12/raizes-do-nordeste-api/src/controller"
	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.RouterGroup) {

	r.POST("/usuarios", controller.CreateUser)
	r.GET("/usuarios/:userId", controller.FindUserById)
	r.GET("/usuarios", controller.FindUserByEmail)
	r.PUT("/usuarios/:userId", controller.UpdateUser)
	r.DELETE("/usuarios/:userId", controller.DeleteUser)

}