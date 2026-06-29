package service

import (
	"github.com/CauaMarques12/raizes-do-nordeste-api/src/configuration/rest_err"
	"github.com/CauaMarques12/raizes-do-nordeste-api/src/model"
	"github.com/CauaMarques12/raizes-do-nordeste-api/src/model/repository"
)

func NewProductDomainService() ProductDomainService {
	return &productDomainService{
		repository: repository.NewProductRepository(),
	}
}

type productDomainService struct {
	repository repository.ProductRepository
}

type ProductDomainService interface {
	CreateProduct(model.ProductDomainInterface) *rest_err.RestErr
	FindProduct(string) (model.ProductDomainInterface, *rest_err.RestErr)
	FindProducts(string, int64, int64) ([]model.ProductDomainInterface, *rest_err.RestErr)
	UpdateProduct(string, model.ProductDomainInterface) (model.ProductDomainInterface, *rest_err.RestErr)
}
