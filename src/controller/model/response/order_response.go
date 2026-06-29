package response

import "time"

type OrderItemResponse struct {
	ProductID      string `json:"produtoId"`
	Quantity       int64  `json:"quantidade"`
	UnitPriceCents int64  `json:"precoUnitarioCents"`
	SubtotalCents  int64  `json:"subtotalCents"`
}

type OrderResponse struct {
	ID            string              `json:"id"`
	ClientID      string              `json:"clienteId"`
	UnitID        string              `json:"unidadeId"`
	Channel       string              `json:"canalPedido"`
	PaymentMethod string              `json:"formaPagamento"`
	PromotionCode string              `json:"codigoPromocao,omitempty"`
	Status        string              `json:"status"`
	TotalCents    int64               `json:"totalCents"`
	DiscountCents int64               `json:"descontoCents"`
	Items         []OrderItemResponse `json:"itens"`
	CreatedAt     time.Time           `json:"createdAt"`
	UpdatedAt     time.Time           `json:"updatedAt"`
}
