package response

import "time"

type UserResponse struct {
	ID                   string    `json:"id"`
	Email                string    `json:"email"`
	Name                 string    `json:"name"`
	Role                 string    `json:"role"`
	FidelidadeConsentida bool      `json:"fidelidadeConsentida"`
	Active               bool      `json:"active"`
	CreatedAt            time.Time `json:"createdAt"`
	UpdatedAt            time.Time `json:"updatedAt"`
}
