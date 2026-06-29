package service

import (
	"github.com/CauaMarques12/raizes-do-nordeste-api/src/configuration/rest_err"
	"github.com/CauaMarques12/raizes-do-nordeste-api/src/model"
	"github.com/CauaMarques12/raizes-do-nordeste-api/src/model/repository"
)

func NewPaymentDomainService() PaymentDomainService {
	return &paymentDomainService{
		paymentRepository: repository.NewPaymentRepository(),
		orderRepository:   repository.NewOrderRepository(),
		userRepository:    repository.NewUserRepository(),
		stockRepository:   repository.NewStockRepository(),
		loyaltyRepository: repository.NewLoyaltyRepository(),
	}
}

type paymentDomainService struct {
	paymentRepository repository.PaymentRepository
	orderRepository   repository.OrderRepository
	userRepository    repository.UserRepository
	stockRepository   repository.StockRepository
	loyaltyRepository repository.LoyaltyRepository
}

type PaymentDomainService interface {
	ProcessPayment(model.PaymentDomainInterface) (model.PaymentDomainInterface, *rest_err.RestErr)
	FindPayment(string) (model.PaymentDomainInterface, *rest_err.RestErr)
	FindPaymentsByOrderID(string, int64, int64) ([]model.PaymentDomainInterface, *rest_err.RestErr)
}
