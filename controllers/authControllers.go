package controllers

import (
	"eduapp-backend/models"
	"eduapp-backend/services"
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserRemote struct {
	UserCalls services.AuthContracts
}

func UserRemConstruct(UserMethods services.AuthContracts) *UserRemote {
	return &UserRemote{
		UserCalls: UserMethods,
	}
}

func (authApp *UserRemote) UserRoutes(incomingRoutes *gin.RouterGroup) {
	router := incomingRoutes.Group("/auth")
	router.POST("/login", authApp.Login)
	router.POST("/signup", authApp.SignUp)
}

func (authApp *UserRemote) Login(c *gin.Context) {

	var loginInst models.User

	if err := c.BindJSON(&loginInst); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"ok":      false,
			"message": "Invalid Creds",
			"error":   err,
		})
		panic(err)
	}
	// fmt.Println(loginInst)

	ok, user, err := authApp.UserCalls.Login(loginInst.Email, loginInst.Password)
	if err != nil {
		c.JSON(500, gin.H{
			"ok":      false,
			"error":   err.Error(),
			"message": "wrong pass",
		})
		panic(err)
	}

	if !ok {
		c.JSON(200, gin.H{
			"ok":      ok,
			"message": "Wrong Pass",
			"error":   err,
		})
		return
	}

	c.JSON(200, gin.H{
		"user":    user,
		"ok":      ok,
		"message": "Log In!",
		"error":   nil,
	})

}

func (authApp *UserRemote) SignUp(c *gin.Context) {

	var signUpInst models.User

	if err := c.BindJSON(&signUpInst); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"ok":      false,
			"message": "Invalid Creds",
			"error":   err,
		})
		panic(err)
	}
	fmt.Println(signUpInst)

	ok, err := authApp.UserCalls.SignUp(signUpInst.UserName, signUpInst.Email, signUpInst.Password)
	if err != nil {
		c.JSON(500, gin.H{
			"error": errors.New("etho"),
		})
		panic(err)
	}

	if !ok {
		c.JSON(200, gin.H{
			"ok":      ok,
			"message": "Wrong Pass",
			"error":   err,
		})
		return
	}
	c.JSON(200, gin.H{
		"ok":      ok,
		"message": "Signed Up!",
		"error":   nil,
	})

}
