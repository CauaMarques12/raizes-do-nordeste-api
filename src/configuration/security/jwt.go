package security

import (
	"os"
	"strconv"
	"time"

	"github.com/CauaMarques12/raizes-do-nordeste-api/src/configuration/rest_err"
	"github.com/CauaMarques12/raizes-do-nordeste-api/src/model"
	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	UserID string `json:"userId"`
	Email  string `json:"email"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}

func GenerateToken(userDomain model.UserDomainInterface) (string, int64, *rest_err.RestErr) {
	expiresIn := getExpiresIn()
	claims := Claims{
		UserID: userDomain.GetID(),
		Email:  userDomain.GetEmail(),
		Role:   userDomain.GetRole(),
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().UTC().Add(time.Duration(expiresIn) * time.Second)),
			IssuedAt:  jwt.NewNumericDate(time.Now().UTC()),
			Subject:   userDomain.GetID(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(getSecret())
	if err != nil {
		return "", 0, rest_err.NewInternalServerError("Error trying to generate token")
	}

	return tokenString, expiresIn, nil
}

func ValidateToken(tokenString string) (*Claims, *rest_err.RestErr) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return getSecret(), nil
	})
	if err != nil || !token.Valid {
		return nil, rest_err.NewUnauthorizedRequestError("Invalid token")
	}

	return claims, nil
}

func getSecret() []byte {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		secret = "raizes-do-nordeste-dev-secret"
	}

	return []byte(secret)
}

func getExpiresIn() int64 {
	expiresIn, err := strconv.ParseInt(os.Getenv("JWT_EXPIRES_IN"), 10, 64)
	if err != nil || expiresIn <= 0 {
		return 3600
	}

	return expiresIn
}
