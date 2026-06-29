package routes

import (
	"net/http"

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
	paymentController controller.PaymentControllerInterface,
	loyaltyController controller.LoyaltyControllerInterface,
	menuController controller.MenuControllerInterface,
	promotionController controller.PromotionControllerInterface,
) {
	r.GET("/swagger", serveSwagger)
	r.GET("/swagger/index.html", serveSwagger)
	r.GET("/swagger/openapi.yaml", func(c *gin.Context) {
		c.Header("Content-Type", "application/yaml; charset=utf-8")
		c.File("docs/openapi.yaml")
	})

	r.POST("/auth/login", authController.Login)
	r.POST("/usuarios", userController.CreateUser)

	auth := r.Group("/")
	auth.Use(middleware.AuthMiddleware())
	auth.GET("/usuarios/me", userController.FindLoggedUser)
	auth.PATCH("/usuarios/me/consentimentos/fidelidade", userController.UpdateLoggedUserFidelityConsent)
	auth.GET("/usuarios/:userId", middleware.RoleMiddleware("ADMIN", "GERENTE"), userController.FindUserById)
	auth.GET("/usuarios", middleware.RoleMiddleware("ADMIN", "GERENTE"), userController.FindUserByEmail)
	auth.PATCH("/usuarios/:userId", middleware.RoleMiddleware("ADMIN", "GERENTE"), userController.UpdateUser)
	auth.DELETE("/usuarios/:userId", middleware.RoleMiddleware("ADMIN"), userController.DeleteUser)

	auth.POST("/unidades", middleware.RoleMiddleware("ADMIN", "GERENTE"), unitController.CreateUnit)
	auth.GET("/unidades", unitController.FindUnits)
	auth.GET("/unidades/:unitId", unitController.FindUnitById)
	auth.GET("/unidades/:unitId/cardapio", menuController.FindMenuByUnit)
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

	auth.POST("/pagamentos", middleware.RoleMiddleware("ADMIN", "GERENTE", "ATENDENTE", "CLIENTE"), paymentController.CreatePayment)
	auth.GET("/pagamentos", middleware.RoleMiddleware("ADMIN", "GERENTE", "ATENDENTE", "CLIENTE"), paymentController.FindPaymentsByOrderId)
	auth.GET("/pagamentos/:paymentId", middleware.RoleMiddleware("ADMIN", "GERENTE", "ATENDENTE", "CLIENTE"), paymentController.FindPaymentById)

	auth.GET("/fidelidade/saldo", middleware.RoleMiddleware("ADMIN", "GERENTE", "CLIENTE"), loyaltyController.FindBalance)
	auth.GET("/fidelidade/historico", middleware.RoleMiddleware("ADMIN", "GERENTE", "CLIENTE"), loyaltyController.FindHistory)
	auth.POST("/fidelidade/resgates", middleware.RoleMiddleware("ADMIN", "GERENTE", "CLIENTE"), loyaltyController.RedeemPoints)

	auth.POST("/promocoes", middleware.RoleMiddleware("ADMIN", "GERENTE"), promotionController.CreatePromotion)
	auth.GET("/promocoes", promotionController.FindPromotions)
	auth.GET("/promocoes/:promotionId", promotionController.FindPromotionById)
	auth.PATCH("/promocoes/:promotionId", middleware.RoleMiddleware("ADMIN", "GERENTE"), promotionController.UpdatePromotion)
}

func serveSwagger(c *gin.Context) {
	c.Data(http.StatusOK, "text/html; charset=utf-8", []byte(`<!doctype html>
<html lang="pt-BR">
<head>
  <meta charset="utf-8">
  <title>Raizes do Nordeste API</title>
  <link rel="stylesheet" href="https://unpkg.com/swagger-ui-dist@5/swagger-ui.css">
</head>
<body>
  <div id="swagger-ui"></div>
  <script src="https://unpkg.com/swagger-ui-dist@5/swagger-ui-bundle.js"></script>
  <script>
    window.onload = function() {
      SwaggerUIBundle({
        url: "/swagger/openapi.yaml",
        dom_id: "#swagger-ui"
      });
    };
  </script>
</body>
</html>`))
}
