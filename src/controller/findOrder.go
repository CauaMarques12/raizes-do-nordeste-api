package controller

import (
	"net/http"

	"github.com/CauaMarques12/raizes-do-nordeste-api/src/configuration/logger"
	"github.com/CauaMarques12/raizes-do-nordeste-api/src/configuration/rest_err"
	"github.com/CauaMarques12/raizes-do-nordeste-api/src/view"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (oc *orderControllerInterface) FindOrderById(c *gin.Context) {
	logger.Info("Iniciando busca de pedido por id", zap.String("jornada", "find_order"))
	orderID := c.Param("orderId")
	orderDomain, err := oc.service.FindOrder(orderID)
	if err != nil {
		logger.Error("Erro ao tentar buscar pedido por id", err, zap.String("jornada", "find_order"))
		c.JSON(err.Code, err)
		return
	}

	c.JSON(http.StatusOK, view.ConvertOrderDomainToResponse(orderDomain))
}

func (oc *orderControllerInterface) FindOrders(c *gin.Context) {
	logger.Info("Iniciando listagem de pedidos", zap.String("jornada", "find_orders"))
	channel := c.Query("canalPedido")
	status := c.Query("status")

	if channel != "" && !isValidOrderChannel(channel) {
		err := rest_err.NewBadRequestError("Invalid canalPedido")
		c.JSON(err.Code, err)
		return
	}

	orderDomains, err := oc.service.FindOrders(channel, status)
	if err != nil {
		logger.Error("Erro ao tentar listar pedidos", err, zap.String("jornada", "find_orders"))
		c.JSON(err.Code, err)
		return
	}

	c.JSON(http.StatusOK, view.ConvertOrderDomainsToResponse(orderDomains))
}

func isValidOrderChannel(channel string) bool {
	validChannels := map[string]bool{
		"APP":    true,
		"TOTEM":  true,
		"BALCAO": true,
		"PICKUP": true,
		"WEB":    true,
	}

	return validChannels[channel]
}
