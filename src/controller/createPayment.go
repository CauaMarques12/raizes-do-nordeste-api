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

func (pc *paymentControllerInterface) CreatePayment(c *gin.Context) {
	logger.Info("Iniciando pagamento mock", zap.String("jornada", "create_payment"))
	var paymentRequest request.PaymentRequest
	if err := c.ShouldBindJSON(&paymentRequest); err != nil {
		logger.Error("Erro ao tentar validar pagamento mock", err, zap.String("jornada", "create_payment"))
		errRest := validation.ValidateUserError(err)
		c.JSON(errRest.Code, errRest)
		return
	}

	domain := model.NewPaymentDomain(
		paymentRequest.OrderID,
		paymentRequest.AmountCents,
		*paymentRequest.Approved,
	)

	paymentDomain, err := pc.service.ProcessPayment(domain)
	if err != nil {
		logger.Error("Erro ao tentar processar pagamento mock", err, zap.String("jornada", "create_payment"))
		c.JSON(err.Code, err)
		return
	}

	logger.Info("Pagamento mock processado com sucesso", zap.String("jornada", "create_payment"))
	c.JSON(http.StatusCreated, view.ConvertPaymentDomainToResponse(paymentDomain))
}
