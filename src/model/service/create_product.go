package service

import (
	"github.com/CauaMarques12/raizes-do-nordeste-api/src/configuration/rest_err"
	"github.com/CauaMarques12/raizes-do-nordeste-api/src/model"
)

func (pd *productDomainService) CreateProduct(productDomain model.ProductDomainInterface) *rest_err.RestErr {
	return pd.repository.CreateProduct(productDomain)
}
