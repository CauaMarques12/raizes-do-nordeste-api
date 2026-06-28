package service

import (
	"github.com/CauaMarques12/raizes-do-nordeste-api/src/configuration/rest_err"
	"github.com/CauaMarques12/raizes-do-nordeste-api/src/configuration/security"
	"github.com/CauaMarques12/raizes-do-nordeste-api/src/model"
	"golang.org/x/crypto/bcrypt"
)

func (ad *authDomainService) Login(
	email, password string,
) (string, int64, model.UserDomainInterface, *rest_err.RestErr) {
	userDomain, err := ad.repository.FindUserByEmail(email)
	if err != nil {
		return "", 0, nil, rest_err.NewUnauthorizedRequestError("Invalid credentials")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(userDomain.GetPassword()), []byte(password)); err != nil {
		return "", 0, nil, rest_err.NewUnauthorizedRequestError("Invalid credentials")
	}

	token, expiresIn, tokenErr := security.GenerateToken(userDomain)
	if tokenErr != nil {
		return "", 0, nil, tokenErr
	}

	return token, expiresIn, userDomain, nil
}
