package service

import (
	"github.com/CauaMarques12/raizes-do-nordeste-api/src/configuration/rest_err"
	"github.com/CauaMarques12/raizes-do-nordeste-api/src/model"
	"github.com/CauaMarques12/raizes-do-nordeste-api/src/model/repository"
)

func NewOrderDomainService() OrderDomainService {
	return &orderDomainService{
		orderRepository:   repository.NewOrderRepository(),
		userRepository:    repository.NewUserRepository(),
		unitRepository:    repository.NewUnitRepository(),
		productRepository: repository.NewProductRepository(),
		stockRepository:   repository.NewStockRepository(),
	}
}

type orderDomainService struct {
	orderRepository   repository.OrderRepository
	userRepository    repository.UserRepository
	unitRepository    repository.UnitRepository
	productRepository repository.ProductRepository
	stockRepository   repository.StockRepository
}

type OrderDomainService interface {
	CreateOrder(model.OrderDomainInterface) *rest_err.RestErr
	FindOrder(string) (model.OrderDomainInterface, *rest_err.RestErr)
	FindOrders(channel, status string) ([]model.OrderDomainInterface, *rest_err.RestErr)
	UpdateStatus(orderID, status string) (model.OrderDomainInterface, *rest_err.RestErr)
	CancelOrder(orderID string) (model.OrderDomainInterface, *rest_err.RestErr)
}
