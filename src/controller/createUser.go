package controller

import (
	"net/http"

	"github.com/CauaMarques12/raizes-do-nordeste-api/src/configuration/logger"
	"github.com/CauaMarques12/raizes-do-nordeste-api/src/configuration/validation"
	"github.com/CauaMarques12/raizes-do-nordeste-api/src/controller/model/request"
	"github.com/CauaMarques12/raizes-do-nordeste-api/src/model"
	"github.com/CauaMarques12/raizes-do-nordeste-api/src/model/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)


var (
    UserDomainInterface model.UserDomainInterface
)

func CreateUser(c *gin.Context) {
   
   logger.Info("Iniciando controlador de criação de usuário", 
	  zap.String("jornada", "create_user"),
)
   var userRequest request.UserRequest

	
   if err := c.ShouldBindJSON(&userRequest); err !=nil {
	logger.Error("Erro ao tentar validar as informações do usuário", err,
	zap.String("jornada", "create_user"))
	errRest := validation.ValidateUserError(err)
	
	c.JSON(errRest.Code, errRest)
	return
   }


   domain := model.NewUserDomain(
	userRequest.Email,
	userRequest.Password,
	userRequest.Name,
	userRequest.Age,
   )
   service := service.NewUserDomainService()
    if err := service.CreateUser(domain); err != nil { 
	c.JSON(err.Code, err)
	return

	}
   
	logger.Info("Usuario criado com sucesso",   zap.String("jornada", "create_user"))
	c.String(http.StatusOK, "")
  
 }