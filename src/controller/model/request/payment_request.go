package request

type PaymentRequest struct {
	OrderID     string `json:"pedidoId" binding:"required"`
	AmountCents int64  `json:"valorCents" binding:"required,min=1"`
	Approved    *bool  `json:"aprovado" binding:"required"`
}
