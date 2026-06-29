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

func (pc *promotionControllerInterface) CreatePromotion(c *gin.Context) {
	logger.Info("Iniciando criacao de promocao", zap.String("jornada", "create_promotion"))
	var promotionRequest request.PromotionRequest
	if err := c.ShouldBindJSON(&promotionRequest); err != nil {
		logger.Error("Erro ao tentar validar promocao", err, zap.String("jornada", "create_promotion"))
		errRest := validation.ValidateUserError(err)
		c.JSON(errRest.Code, errRest)
		return
	}

	domain := model.NewPromotionDomain(
		promotionRequest.Name,
		promotionRequest.Code,
		promotionRequest.DiscountPercent,
		promotionRequest.Active,
	)

	if err := pc.service.CreatePromotion(domain); err != nil {
		logger.Error("Erro ao tentar criar promocao", err, zap.String("jornada", "create_promotion"))
		c.JSON(err.Code, err)
		return
	}

	logger.Info("Promocao criada com sucesso", zap.String("jornada", "create_promotion"))
	c.JSON(http.StatusCreated, view.ConvertPromotionDomainToResponse(domain))
}
