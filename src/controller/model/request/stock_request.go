package request

type StockMovementRequest struct {
	UnitID    string `json:"unidadeId" binding:"required"`
	ProductID string `json:"produtoId" binding:"required"`
	Type      string `json:"tipo" binding:"required,oneof=ENTRADA SAIDA"`
	Quantity  int64  `json:"quantidade" binding:"required,min=1"`
	Reason    string `json:"motivo" binding:"required,min=3,max=200"`
}
