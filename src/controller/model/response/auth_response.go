package response

type AuthResponse struct {
	AccessToken string       `json:"accessToken"`
	TokenType   string       `json:"tokenType"`
	ExpiresIn   int64        `json:"expiresIn"`
	User        UserResponse `json:"user"`
}
