package response

type MenuItemResponse struct {
	ProductID   string `json:"produtoId"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Category    string `json:"category"`
	PriceCents  int64  `json:"priceCents"`
	Quantity    int64  `json:"quantidade"`
	Available   bool   `json:"disponivel"`
}
