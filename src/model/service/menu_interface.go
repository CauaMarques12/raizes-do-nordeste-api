package service

import (
	"github.com/CauaMarques12/raizes-do-nordeste-api/src/configuration/rest_err"
	"github.com/CauaMarques12/raizes-do-nordeste-api/src/model"
	"github.com/CauaMarques12/raizes-do-nordeste-api/src/model/repository"
)

func NewMenuDomainService() MenuDomainService {
	return &menuDomainService{
		unitRepository:    repository.NewUnitRepository(),
		productRepository: repository.NewProductRepository(),
		stockRepository:   repository.NewStockRepository(),
	}
}

type menuDomainService struct {
	unitRepository    repository.UnitRepository
	productRepository repository.ProductRepository
	stockRepository   repository.StockRepository
}

type MenuDomainService interface {
	FindMenuByUnit(unitID, category string, page, limit int64) ([]model.MenuItemDomainInterface, *rest_err.RestErr)
}
