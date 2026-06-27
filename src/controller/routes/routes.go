package routes

import (
	"github.com/CauaMarques12/raizes-do-nordeste-api/src/controller"
	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.RouterGroup,
	userController controller.UserControllerInterface) {

	r.POST("/usuarios", userController.CreateUser)
	r.GET("/usuarios/:userId", userController.FindUserById)
	r.GET("/usuarios", userController.FindUserByEmail)
	r.PATCH("/usuarios/:userId", userController.UpdateUser)
	r.DELETE("/usuarios/:userId", userController.DeleteUser)

}
