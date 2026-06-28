package request

type UnitRequest struct {
	Name    string `json:"name" binding:"required,min=3,max=100"`
	Address string `json:"address" binding:"required,min=5,max=200"`
	City    string `json:"city" binding:"required,min=2,max=100"`
	State   string `json:"state" binding:"required,len=2"`
}

type UnitUpdateRequest struct {
	Name    string `json:"name" binding:"omitempty,min=3,max=100"`
	Address string `json:"address" binding:"omitempty,min=5,max=200"`
	City    string `json:"city" binding:"omitempty,min=2,max=100"`
	State   string `json:"state" binding:"omitempty,len=2"`
}
