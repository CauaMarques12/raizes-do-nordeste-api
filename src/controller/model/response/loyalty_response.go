package response

import "time"

type LoyaltyBalanceResponse struct {
	ID        string    `json:"id"`
	ClientID  string    `json:"clienteId"`
	Points    int64     `json:"pontos"`
	Active    bool      `json:"active"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type LoyaltyMovementResponse struct {
	ID           string    `json:"id"`
	ClientID     string    `json:"clienteId"`
	Type         string    `json:"tipo"`
	Points       int64     `json:"pontos"`
	Reason       string    `json:"motivo"`
	OrderID      string    `json:"pedidoId,omitempty"`
	BalanceAfter int64     `json:"saldoAposMovimentacao"`
	CreatedAt    time.Time `json:"createdAt"`
}
