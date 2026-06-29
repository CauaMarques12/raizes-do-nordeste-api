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

func (pc *promotionControllerInterface) UpdatePromotion(c *gin.Context) {
	logger.Info("Iniciando atualizacao de promocao", zap.String("jornada", "update_promotion"))
	promotionID := c.Param("promotionId")
	var promotionRequest request.PromotionUpdateRequest
	if err := c.ShouldBindJSON(&promotionRequest); err != nil {
		logger.Error("Erro ao tentar validar atualizacao de promocao", err, zap.String("jornada", "update_promotion"))
		errRest := validation.ValidateUserError(err)
		c.JSON(errRest.Code, errRest)
		return
	}

	domain := model.NewPromotionUpdateDomain(
		promotionRequest.Name,
		promotionRequest.Code,
		promotionRequest.DiscountPercent,
		promotionRequest.Active,
	)

	updatedPromotion, err := pc.service.UpdatePromotion(promotionID, domain)
	if err != nil {
		logger.Error("Erro ao tentar atualizar promocao", err, zap.String("jornada", "update_promotion"))
		c.JSON(err.Code, err)
		return
	}

	logger.Info("Promocao atualizada com sucesso", zap.String("jornada", "update_promotion"))
	c.JSON(http.StatusOK, view.ConvertPromotionDomainToResponse(updatedPromotion))
}
