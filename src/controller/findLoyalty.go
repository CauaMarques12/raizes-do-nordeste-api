package controller

import (
	"net/http"

	"github.com/CauaMarques12/raizes-do-nordeste-api/src/configuration/logger"
	"github.com/CauaMarques12/raizes-do-nordeste-api/src/configuration/rest_err"
	"github.com/CauaMarques12/raizes-do-nordeste-api/src/view"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (lc *loyaltyControllerInterface) FindBalance(c *gin.Context) {
	logger.Info("Iniciando consulta de saldo de fidelidade", zap.String("jornada", "find_loyalty_balance"))
	clientID, err := resolveLoyaltyClientID(c, c.Query("clienteId"))
	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	balanceDomain, restErr := lc.service.FindBalance(clientID)
	if restErr != nil {
		logger.Error("Erro ao tentar consultar saldo de fidelidade", restErr, zap.String("jornada", "find_loyalty_balance"))
		c.JSON(restErr.Code, restErr)
		return
	}

	c.JSON(http.StatusOK, view.ConvertLoyaltyBalanceDomainToResponse(balanceDomain))
}

func (lc *loyaltyControllerInterface) FindHistory(c *gin.Context) {
	logger.Info("Iniciando consulta de historico de fidelidade", zap.String("jornada", "find_loyalty_history"))
	clientID, err := resolveLoyaltyClientID(c, c.Query("clienteId"))
	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	movementsDomain, restErr := lc.service.FindMovements(clientID)
	if restErr != nil {
		logger.Error("Erro ao tentar consultar historico de fidelidade", restErr, zap.String("jornada", "find_loyalty_history"))
		c.JSON(restErr.Code, restErr)
		return
	}

	c.JSON(http.StatusOK, view.ConvertLoyaltyMovementDomainsToResponse(movementsDomain))
}

func resolveLoyaltyClientID(c *gin.Context, requestedClientID string) (string, *rest_err.RestErr) {
	userID := c.GetString("userId")
	userRole := c.GetString("userRole")
	if requestedClientID == "" {
		requestedClientID = userID
	}

	if userRole == "CLIENTE" && requestedClientID != userID {
		return "", rest_err.NewForbiddenError("Usuario nao pode acessar fidelidade de outro cliente")
	}

	return requestedClientID, nil
}
