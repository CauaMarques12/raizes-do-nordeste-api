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

func ConvertStockBalanceDomainToResponse(
	stockDomain model.StockBalanceDomainInterface,
) response.StockBalanceResponse {
	return response.StockBalanceResponse{
		ID:        stockDomain.GetID(),
		UnitID:    stockDomain.GetUnitID(),
		ProductID: stockDomain.GetProductID(),
		Quantity:  stockDomain.GetQuantity(),
		Active:    stockDomain.GetActive(),
		CreatedAt: stockDomain.GetCreatedAt(),
		UpdatedAt: stockDomain.GetUpdatedAt(),
	}
}

func ConvertStockMovementDomainToResponse(
	stockMovementDomain model.StockMovementDomainInterface,
) response.StockMovementResponse {
	return response.StockMovementResponse{
		ID:           stockMovementDomain.GetID(),
		UnitID:       stockMovementDomain.GetUnitID(),
		ProductID:    stockMovementDomain.GetProductID(),
		Type:         stockMovementDomain.GetType(),
		Quantity:     stockMovementDomain.GetQuantity(),
		Reason:       stockMovementDomain.GetReason(),
		BalanceAfter: stockMovementDomain.GetBalanceAfter(),
		CreatedAt:    stockMovementDomain.GetCreatedAt(),
	}
}

func ConvertOrderDomainToResponse(
	orderDomain model.OrderDomainInterface,
) response.OrderResponse {
	return response.OrderResponse{
		ID:            orderDomain.GetID(),
		ClientID:      orderDomain.GetClientID(),
		UnitID:        orderDomain.GetUnitID(),
		Channel:       orderDomain.GetChannel(),
		PaymentMethod: orderDomain.GetPaymentMethod(),
		Status:        orderDomain.GetStatus(),
		TotalCents:    orderDomain.GetTotalCents(),
		Items:         convertOrderItemsToResponse(orderDomain.GetItems()),
		CreatedAt:     orderDomain.GetCreatedAt(),
		UpdatedAt:     orderDomain.GetUpdatedAt(),
	}
}

func ConvertOrderDomainsToResponse(
	orderDomains []model.OrderDomainInterface,
) []response.OrderResponse {
	orderResponses := make([]response.OrderResponse, 0, len(orderDomains))
	for _, orderDomain := range orderDomains {
		orderResponses = append(orderResponses, ConvertOrderDomainToResponse(orderDomain))
	}

	return orderResponses
}

func ConvertLoyaltyBalanceDomainToResponse(
	loyaltyBalanceDomain model.LoyaltyBalanceDomainInterface,
) response.LoyaltyBalanceResponse {
	return response.LoyaltyBalanceResponse{
		ID:        loyaltyBalanceDomain.GetID(),
		ClientID:  loyaltyBalanceDomain.GetUserID(),
		Points:    loyaltyBalanceDomain.GetPoints(),
		Active:    loyaltyBalanceDomain.GetActive(),
		CreatedAt: loyaltyBalanceDomain.GetCreatedAt(),
		UpdatedAt: loyaltyBalanceDomain.GetUpdatedAt(),
	}
}

func ConvertLoyaltyMovementDomainToResponse(
	loyaltyMovementDomain model.LoyaltyMovementDomainInterface,
) response.LoyaltyMovementResponse {
	return response.LoyaltyMovementResponse{
		ID:           loyaltyMovementDomain.GetID(),
		ClientID:     loyaltyMovementDomain.GetUserID(),
		Type:         loyaltyMovementDomain.GetType(),
		Points:       loyaltyMovementDomain.GetPoints(),
		Reason:       loyaltyMovementDomain.GetReason(),
		OrderID:      loyaltyMovementDomain.GetOrderID(),
		BalanceAfter: loyaltyMovementDomain.GetBalanceAfter(),
		CreatedAt:    loyaltyMovementDomain.GetCreatedAt(),
	}
}

func ConvertLoyaltyMovementDomainsToResponse(
	loyaltyMovementDomains []model.LoyaltyMovementDomainInterface,
) []response.LoyaltyMovementResponse {
	loyaltyMovementResponses := make([]response.LoyaltyMovementResponse, 0, len(loyaltyMovementDomains))
	for _, loyaltyMovementDomain := range loyaltyMovementDomains {
		loyaltyMovementResponses = append(loyaltyMovementResponses, ConvertLoyaltyMovementDomainToResponse(loyaltyMovementDomain))
	}

	return loyaltyMovementResponses
}

func ConvertPaymentDomainToResponse(
	paymentDomain model.PaymentDomainInterface,
) response.PaymentResponse {
	return response.PaymentResponse{
		ID:                   paymentDomain.GetID(),
		OrderID:              paymentDomain.GetOrderID(),
		Method:               paymentDomain.GetMethod(),
		AmountCents:          paymentDomain.GetAmountCents(),
		Status:               paymentDomain.GetStatus(),
		GatewayTransactionID: paymentDomain.GetGatewayTransactionID(),
		Message:              paymentDomain.GetMessage(),
		CreatedAt:            paymentDomain.GetCreatedAt(),
		UpdatedAt:            paymentDomain.GetUpdatedAt(),
	}
}

func ConvertPaymentDomainsToResponse(
	paymentDomains []model.PaymentDomainInterface,
) []response.PaymentResponse {
	paymentResponses := make([]response.PaymentResponse, 0, len(paymentDomains))
	for _, paymentDomain := range paymentDomains {
		paymentResponses = append(paymentResponses, ConvertPaymentDomainToResponse(paymentDomain))
	}

	return paymentResponses
}

func convertOrderItemsToResponse(
	items []model.OrderItemDomain,
) []response.OrderItemResponse {
	itemResponses := make([]response.OrderItemResponse, 0, len(items))
	for _, item := range items {
		itemResponses = append(itemResponses, response.OrderItemResponse{
			ProductID:      item.ProductID,
			Quantity:       item.Quantity,
			UnitPriceCents: item.UnitPriceCents,
			SubtotalCents:  item.SubtotalCents,
		})
	}

	return itemResponses
}
