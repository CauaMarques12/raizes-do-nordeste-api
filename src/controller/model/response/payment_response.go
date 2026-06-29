package response

import "time"

type PaymentResponse struct {
	ID                   string    `json:"id"`
	OrderID              string    `json:"pedidoId"`
	Method               string    `json:"metodo"`
	AmountCents          int64     `json:"valorCents"`
	Status               string    `json:"status"`
	GatewayTransactionID string    `json:"gatewayTransactionId"`
	Message              string    `json:"mensagem"`
	CreatedAt            time.Time `json:"createdAt"`
	UpdatedAt            time.Time `json:"updatedAt"`
}
