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
	unitController controller.UnitControllerInterface,
	productController controller.ProductControllerInterface,
	stockController controller.StockControllerInterface,
	orderController controller.OrderControllerInterface,
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

	auth.POST("/unidades", middleware.RoleMiddleware("ADMIN", "GERENTE"), unitController.CreateUnit)
	auth.GET("/unidades", unitController.FindUnits)
	auth.GET("/unidades/:unitId", unitController.FindUnitById)
	auth.PATCH("/unidades/:unitId", middleware.RoleMiddleware("ADMIN", "GERENTE"), unitController.UpdateUnit)

	auth.POST("/produtos", middleware.RoleMiddleware("ADMIN", "GERENTE"), productController.CreateProduct)
	auth.GET("/produtos", productController.FindProducts)
	auth.GET("/produtos/:productId", productController.FindProductById)
	auth.PATCH("/produtos/:productId", middleware.RoleMiddleware("ADMIN", "GERENTE"), productController.UpdateProduct)

	auth.POST("/estoque/movimentacoes", middleware.RoleMiddleware("ADMIN", "GERENTE"), stockController.CreateStockMovement)
	auth.GET("/estoque", middleware.RoleMiddleware("ADMIN", "GERENTE", "ATENDENTE"), stockController.FindStockBalance)

	auth.POST("/pedidos", middleware.RoleMiddleware("ADMIN", "GERENTE", "ATENDENTE", "CLIENTE"), orderController.CreateOrder)
	auth.GET("/pedidos", middleware.RoleMiddleware("ADMIN", "GERENTE", "ATENDENTE", "COZINHA"), orderController.FindOrders)
	auth.GET("/pedidos/:orderId", middleware.RoleMiddleware("ADMIN", "GERENTE", "ATENDENTE", "COZINHA", "CLIENTE"), orderController.FindOrderById)
	auth.PATCH("/pedidos/:orderId/status", middleware.RoleMiddleware("ADMIN", "GERENTE", "ATENDENTE", "COZINHA"), orderController.UpdateOrderStatus)
	auth.PATCH("/pedidos/:orderId/cancelamento", middleware.RoleMiddleware("ADMIN", "GERENTE", "ATENDENTE", "CLIENTE"), orderController.CancelOrder)
}
