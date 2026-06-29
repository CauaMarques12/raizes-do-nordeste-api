package controller

import (
	"net/http"

	"github.com/CauaMarques12/raizes-do-nordeste-api/src/configuration/logger"
	"github.com/CauaMarques12/raizes-do-nordeste-api/src/configuration/validation"
	"github.com/CauaMarques12/raizes-do-nordeste-api/src/controller/model/request"
	"github.com/CauaMarques12/raizes-do-nordeste-api/src/model"
	"github.com/CauaMarques12/raizes-do-nordeste-api/src/view"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (oc *orderControllerInterface) CreateOrder(c *gin.Context) {
	logger.Info("Iniciando criacao de pedido", zap.String("jornada", "create_order"))
	var orderRequest request.OrderRequest
	if err := c.ShouldBindJSON(&orderRequest); err != nil {
		logger.Error("Erro ao tentar validar pedido", err, zap.String("jornada", "create_order"))
		errRest := validation.ValidateUserError(err)
		c.JSON(errRest.Code, errRest)
		return
	}

	domain := model.NewOrderDomain(
		orderRequest.ClientID,
		orderRequest.UnitID,
		orderRequest.Channel,
		orderRequest.PaymentMethod,
		convertOrderItemsRequestToDomain(orderRequest.Items),
	)

	if err := oc.service.CreateOrder(domain); err != nil {
		logger.Error("Erro ao tentar criar pedido", err, zap.String("jornada", "create_order"))
		c.JSON(err.Code, err)
		return
	}

	logger.Info("Pedido criado com sucesso", zap.String("jornada", "create_order"))
	c.JSON(http.StatusCreated, view.ConvertOrderDomainToResponse(domain))
}

func convertOrderItemsRequestToDomain(items []request.OrderItemRequest) []model.OrderItemDomain {
	orderItems := make([]model.OrderItemDomain, 0, len(items))
	for _, item := range items {
		orderItems = append(orderItems, model.OrderItemDomain{
			ProductID: item.ProductID,
			Quantity:  item.Quantity,
		})
	}

	return orderItems
}
