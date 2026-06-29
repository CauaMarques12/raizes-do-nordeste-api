package controller

import (
	"net/http"
	"strconv"

	"github.com/CauaMarques12/raizes-do-nordeste-api/src/configuration/logger"
	"github.com/CauaMarques12/raizes-do-nordeste-api/src/configuration/rest_err"
	"github.com/CauaMarques12/raizes-do-nordeste-api/src/view"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (pc *promotionControllerInterface) FindPromotionById(c *gin.Context) {
	logger.Info("Iniciando busca de promocao por id", zap.String("jornada", "find_promotion"))
	promotionID := c.Param("promotionId")
	promotionDomain, err := pc.service.FindPromotion(promotionID)
	if err != nil {
		logger.Error("Erro ao tentar buscar promocao por id", err, zap.String("jornada", "find_promotion"))
		c.JSON(err.Code, err)
		return
	}

	c.JSON(http.StatusOK, view.ConvertPromotionDomainToResponse(promotionDomain))
}

func (pc *promotionControllerInterface) FindPromotions(c *gin.Context) {
	logger.Info("Iniciando listagem de promocoes", zap.String("jornada", "find_promotions"))
	active, err := getActiveFilter(c.Query("active"))
	if err != nil {
		c.JSON(err.Code, err)
		return
	}
	page, limit, paginationErr := getPagination(c)
	if paginationErr != nil {
		c.JSON(paginationErr.Code, paginationErr)
		return
	}

	promotionDomains, restErr := pc.service.FindPromotions(active, page, limit)
	if restErr != nil {
		logger.Error("Erro ao tentar listar promocoes", restErr, zap.String("jornada", "find_promotions"))
		c.JSON(restErr.Code, restErr)
		return
	}

	c.JSON(http.StatusOK, view.ConvertPromotionDomainsToResponse(promotionDomains))
}

func getActiveFilter(activeQuery string) (*bool, *rest_err.RestErr) {
	if activeQuery == "" {
		return nil, nil
	}

	active, err := strconv.ParseBool(activeQuery)
	if err != nil {
		return nil, rest_err.NewBadRequestError("active deve ser true ou false")
	}

	return &active, nil
}
