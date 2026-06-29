package request

type PromotionRequest struct {
	Name            string `json:"name" binding:"required,min=3,max=100"`
	Code            string `json:"code" binding:"required,alphanum,min=3,max=30"`
	DiscountPercent int64  `json:"discountPercent" binding:"required,min=1,max=90"`
	Active          *bool  `json:"active"`
}

type PromotionUpdateRequest struct {
	Name            string `json:"name" binding:"omitempty,min=3,max=100"`
	Code            string `json:"code" binding:"omitempty,alphanum,min=3,max=30"`
	DiscountPercent int64  `json:"discountPercent" binding:"omitempty,min=1,max=90"`
	Active          *bool  `json:"active"`
}
