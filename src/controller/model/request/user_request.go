package request

type UserRequest struct {
	Email                string `json:"email" binding:"required,email"`
	Password             string `json:"password" binding:"required,min=6,containsany=!@#$%^&*"`
	Name                 string `json:"name" binding:"required,min=4,max=100"`
	Role                 string `json:"role" binding:"omitempty,oneof=CLIENTE ATENDENTE COZINHA GERENTE ADMIN"`
	FidelidadeConsentida *bool  `json:"fidelidadeConsentida"`
}

type UserUpdateRequest struct {
	Email                string `json:"email" binding:"omitempty,email"`
	Password             string `json:"password" binding:"omitempty,min=6,containsany=!@#$%^&*"`
	Name                 string `json:"name" binding:"omitempty,min=4,max=100"`
	Role                 string `json:"role" binding:"omitempty,oneof=CLIENTE ATENDENTE COZINHA GERENTE ADMIN"`
	FidelidadeConsentida *bool  `json:"fidelidadeConsentida"`
}

type FidelityConsentRequest struct {
	FidelidadeConsentida *bool `json:"fidelidadeConsentida" binding:"required"`
}
