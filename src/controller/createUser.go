package controller

import (
	"fmt"
	"log"

	
	"github.com/CauaMarques12/raizes-do-nordeste-api/src/configuration/validation"
	"github.com/CauaMarques12/raizes-do-nordeste-api/src/controller/model/request"
	
	"github.com/gin-gonic/gin"
)


func CreateUser(c *gin.Context) {
   
   log.Println("Init CreateUser Controller")
   var userRequest request.UserRequest


   if err := c.ShouldBindJSON(&userRequest); err !=nil {
	log.Printf("Error trying to marshal object, error=%s", err.Error())
	errRest := validation.ValidateUserError(err)
	
	c.JSON(errRest.Code, errRest)
	return
   }


   fmt.Println(userRequest)
  
 }