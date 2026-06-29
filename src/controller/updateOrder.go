package controller

import (
	"net/http"

	"github.com/CauaMarques12/raizes-do-nordeste-api/src/configuration/logger"
	"github.com/CauaMarques12/raizes-do-nordeste-api/src/configuration/validation"
	"github.com/CauaMarques12/raizes-do-nordeste-api/src/controller/model/request"
	"github.com/CauaMarques12/raizes-do-nordeste-api/src/view"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (oc *orderControllerInterface) UpdateOrderStatus(c *gin.Context) {
	logger.Info("Iniciando atualizacao de status do pedido", zap.String("jornada", "update_order_status"))
	orderID := c.Param("orderId")
	var orderRequest request.OrderStatusRequest
	if err := c.ShouldBindJSON(&orderRequest); err != nil {
		logger.Error("Erro ao tentar validar status do pedido", err, zap.String("jornada", "update_order_status"))
		errRest := validation.ValidateUserError(err)
		c.JSON(errRest.Code, errRest)
		return
	}

	orderDomain, err := oc.service.UpdateStatus(orderID, orderRequest.Status)
	if err != nil {
		logger.Error("Erro ao tentar atualizar status do pedido", err, zap.String("jornada", "update_order_status"))
		c.JSON(err.Code, err)
		return
	}

	logger.Info("Status do pedido atualizado com sucesso", zap.String("jornada", "update_order_status"))
	c.JSON(http.StatusOK, view.ConvertOrderDomainToResponse(orderDomain))
}

func (oc *orderControllerInterface) CancelOrder(c *gin.Context) {
	logger.Info("Iniciando cancelamento de pedido", zap.String("jornada", "cancel_order"))
	orderID := c.Param("orderId")
	orderDomain, err := oc.service.CancelOrder(orderID)
	if err != nil {
		logger.Error("Erro ao tentar cancelar pedido", err, zap.String("jornada", "cancel_order"))
		c.JSON(err.Code, err)
		return
	}

	logger.Info("Pedido cancelado com sucesso", zap.String("jornada", "cancel_order"))
	c.JSON(http.StatusOK, view.ConvertOrderDomainToResponse(orderDomain))
}
