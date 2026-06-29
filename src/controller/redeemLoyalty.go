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

func (lc *loyaltyControllerInterface) RedeemPoints(c *gin.Context) {
	logger.Info("Iniciando resgate de pontos de fidelidade", zap.String("jornada", "redeem_loyalty"))
	var redeemRequest request.LoyaltyRedeemRequest
	if err := c.ShouldBindJSON(&redeemRequest); err != nil {
		logger.Error("Erro ao tentar validar resgate de fidelidade", err, zap.String("jornada", "redeem_loyalty"))
		errRest := validation.ValidateUserError(err)
		c.JSON(errRest.Code, errRest)
		return
	}

	clientID, err := resolveLoyaltyClientID(c, redeemRequest.ClientID)
	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	movementDomain, restErr := lc.service.RedeemPoints(clientID, redeemRequest.Points)
	if restErr != nil {
		logger.Error("Erro ao tentar resgatar pontos de fidelidade", restErr, zap.String("jornada", "redeem_loyalty"))
		c.JSON(restErr.Code, restErr)
		return
	}

	logger.Info("Pontos de fidelidade resgatados com sucesso", zap.String("jornada", "redeem_loyalty"))
	c.JSON(http.StatusCreated, view.ConvertLoyaltyMovementDomainToResponse(movementDomain))
}
