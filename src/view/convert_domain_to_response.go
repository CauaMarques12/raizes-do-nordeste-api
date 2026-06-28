package view

import (
	"github.com/CauaMarques12/raizes-do-nordeste-api/src/controller/model/response"
	"github.com/CauaMarques12/raizes-do-nordeste-api/src/model"
)

func ConvertDomainToResponse(
	userDomain model.UserDomainInterface,
) response.UserResponse {
	return response.UserResponse{
		ID:                   userDomain.GetID(),
		Email:                userDomain.GetEmail(),
		Name:                 userDomain.GetName(),
		Role:                 userDomain.GetRole(),
		FidelidadeConsentida: userDomain.GetFidelidadeConsentida(),
		Active:               userDomain.GetActive(),
		CreatedAt:            userDomain.GetCreatedAt(),
		UpdatedAt:            userDomain.GetUpdatedAt(),
	}
}

func ConvertUnitDomainToResponse(
	unitDomain model.UnitDomainInterface,
) response.UnitResponse {
	return response.UnitResponse{
		ID:        unitDomain.GetID(),
		Name:      unitDomain.GetName(),
		Address:   unitDomain.GetAddress(),
		City:      unitDomain.GetCity(),
		State:     unitDomain.GetState(),
		Active:    unitDomain.GetActive(),
		CreatedAt: unitDomain.GetCreatedAt(),
		UpdatedAt: unitDomain.GetUpdatedAt(),
	}
}

func ConvertUnitDomainsToResponse(
	unitDomains []model.UnitDomainInterface,
) []response.UnitResponse {
	unitResponses := make([]response.UnitResponse, 0, len(unitDomains))
	for _, unitDomain := range unitDomains {
		unitResponses = append(unitResponses, ConvertUnitDomainToResponse(unitDomain))
	}

	return unitResponses
}

func ConvertProductDomainToResponse(
	productDomain model.ProductDomainInterface,
) response.ProductResponse {
	return response.ProductResponse{
		ID:          productDomain.GetID(),
		Name:        productDomain.GetName(),
		Description: productDomain.GetDescription(),
		Category:    productDomain.GetCategory(),
		PriceCents:  productDomain.GetPriceCents(),
		Active:      productDomain.GetActive(),
		CreatedAt:   productDomain.GetCreatedAt(),
		UpdatedAt:   productDomain.GetUpdatedAt(),
	}
}

func ConvertProductDomainsToResponse(
	productDomains []model.ProductDomainInterface,
) []response.ProductResponse {
	productResponses := make([]response.ProductResponse, 0, len(productDomains))
	for _, productDomain := range productDomains {
		productResponses = append(productResponses, ConvertProductDomainToResponse(productDomain))
	}

	return productResponses
}
