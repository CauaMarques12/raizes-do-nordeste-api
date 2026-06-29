package response

import "time"

type StockBalanceResponse struct {
	ID        string    `json:"id"`
	UnitID    string    `json:"unidadeId"`
	ProductID string    `json:"produtoId"`
	Quantity  int64     `json:"quantidade"`
	Active    bool      `json:"active"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type StockMovementResponse struct {
	ID           string    `json:"id"`
	UnitID       string    `json:"unidadeId"`
	ProductID    string    `json:"produtoId"`
	Type         string    `json:"tipo"`
	Quantity     int64     `json:"quantidade"`
	Reason       string    `json:"motivo"`
	BalanceAfter int64     `json:"saldoAposMovimentacao"`
	CreatedAt    time.Time `json:"createdAt"`
}
