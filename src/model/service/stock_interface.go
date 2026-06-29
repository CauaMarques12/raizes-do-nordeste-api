package service

import (
	"github.com/CauaMarques12/raizes-do-nordeste-api/src/configuration/rest_err"
	"github.com/CauaMarques12/raizes-do-nordeste-api/src/model"
	"github.com/CauaMarques12/raizes-do-nordeste-api/src/model/repository"
)

func NewStockDomainService() StockDomainService {
	return &stockDomainService{
		stockRepository:   repository.NewStockRepository(),
		unitRepository:    repository.NewUnitRepository(),
		productRepository: repository.NewProductRepository(),
	}
}

type stockDomainService struct {
	stockRepository   repository.StockRepository
	unitRepository    repository.UnitRepository
	productRepository repository.ProductRepository
}

type StockDomainService interface {
	CreateMovement(model.StockMovementDomainInterface) *rest_err.RestErr
	FindBalance(unitID, productID string) (model.StockBalanceDomainInterface, *rest_err.RestErr)
}
