package service

import (
	"github.com/CauaMarques12/raizes-do-nordeste-api/src/configuration/rest_err"
	"github.com/CauaMarques12/raizes-do-nordeste-api/src/model"
)

func (pd *productDomainService) FindProduct(productID string) (model.ProductDomainInterface, *rest_err.RestErr) {
	return pd.repository.FindProductByID(productID)
}

func (pd *productDomainService) FindProducts(category string) ([]model.ProductDomainInterface, *rest_err.RestErr) {
	return pd.repository.FindProducts(category)
}
