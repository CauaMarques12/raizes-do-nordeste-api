package controller

import (
	"net/http"

	"github.com/CauaMarques12/raizes-do-nordeste-api/src/configuration/logger"
	"github.com/CauaMarques12/raizes-do-nordeste-api/src/configuration/rest_err"
	"github.com/CauaMarques12/raizes-do-nordeste-api/src/view"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (pc *paymentControllerInterface) FindPaymentById(c *gin.Context) {
	logger.Info("Iniciando busca de pagamento por id", zap.String("jornada", "find_payment"))
	paymentID := c.Param("paymentId")
	paymentDomain, err := pc.service.FindPayment(paymentID)
	if err != nil {
		logger.Error("Erro ao tentar buscar pagamento por id", err, zap.String("jornada", "find_payment"))
		c.JSON(err.Code, err)
		return
	}

	c.JSON(http.StatusOK, view.ConvertPaymentDomainToResponse(paymentDomain))
}

func (pc *paymentControllerInterface) FindPaymentsByOrderId(c *gin.Context) {
	logger.Info("Iniciando listagem de pagamentos por pedido", zap.String("jornada", "find_payments"))
	orderID := c.Query("pedidoId")
	if orderID == "" {
		err := rest_err.NewBadRequestError("pedidoId e obrigatorio")
		c.JSON(err.Code, err)
		return
	}

	paymentDomains, err := pc.service.FindPaymentsByOrderID(orderID)
	if err != nil {
		logger.Error("Erro ao tentar listar pagamentos por pedido", err, zap.String("jornada", "find_payments"))
		c.JSON(err.Code, err)
		return
	}

	c.JSON(http.StatusOK, view.ConvertPaymentDomainsToResponse(paymentDomains))
}
