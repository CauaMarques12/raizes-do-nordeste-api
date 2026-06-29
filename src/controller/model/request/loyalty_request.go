package request

type LoyaltyRedeemRequest struct {
	ClientID string `json:"clienteId" binding:"omitempty"`
	Points   int64  `json:"pontos" binding:"required,min=1"`
}
