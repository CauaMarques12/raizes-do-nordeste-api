package request

type OrderItemRequest struct {
	ProductID string `json:"produtoId" binding:"required"`
	Quantity  int64  `json:"quantidade" binding:"required,min=1"`
}

type OrderRequest struct {
	Channel       string             `json:"canalPedido" binding:"required,oneof=APP TOTEM BALCAO PICKUP WEB"`
	ClientID      string             `json:"clienteId" binding:"required"`
	UnitID        string             `json:"unidadeId" binding:"required"`
	Items         []OrderItemRequest `json:"itens" binding:"required,min=1,dive"`
	PaymentMethod string             `json:"formaPagamento" binding:"required,oneof=MOCK"`
}

type OrderStatusRequest struct {
	Status string `json:"status" binding:"required,oneof=AGUARDANDO_PAGAMENTO PAGO EM_PREPARO PRONTO ENTREGUE CANCELADO"`
}
