package request

type ProductRequest struct {
	Name        string `json:"name" binding:"required,min=3,max=100"`
	Description string `json:"description" binding:"required,min=3,max=300"`
	Category    string `json:"category" binding:"required,min=3,max=80"`
	PriceCents  int64  `json:"priceCents" binding:"required,min=1"`
}

type ProductUpdateRequest struct {
	Name        string `json:"name" binding:"omitempty,min=3,max=100"`
	Description string `json:"description" binding:"omitempty,min=3,max=300"`
	Category    string `json:"category" binding:"omitempty,min=3,max=80"`
	PriceCents  int64  `json:"priceCents" binding:"omitempty,min=1"`
}
