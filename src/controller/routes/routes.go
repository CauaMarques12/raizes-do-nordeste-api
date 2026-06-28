package routes

import (
	"github.com/CauaMarques12/raizes-do-nordeste-api/src/configuration/middleware"
	"github.com/CauaMarques12/raizes-do-nordeste-api/src/controller"
	"github.com/gin-gonic/gin"
)

func InitRoutes(
	r *gin.RouterGroup,
	userController controller.UserControllerInterface,
	authController controller.AuthControllerInterface,
) {
	r.POST("/auth/login", authController.Login)
	r.POST("/usuarios", userController.CreateUser)

	auth := r.Group("/")
	auth.Use(middleware.AuthMiddleware())
	auth.GET("/usuarios/me", userController.FindLoggedUser)
	auth.GET("/usuarios/:userId", middleware.RoleMiddleware("ADMIN", "GERENTE"), userController.FindUserById)
	auth.GET("/usuarios", middleware.RoleMiddleware("ADMIN", "GERENTE"), userController.FindUserByEmail)
	auth.PATCH("/usuarios/:userId", middleware.RoleMiddleware("ADMIN", "GERENTE"), userController.UpdateUser)
	auth.DELETE("/usuarios/:userId", middleware.RoleMiddleware("ADMIN"), userController.DeleteUser)
}
