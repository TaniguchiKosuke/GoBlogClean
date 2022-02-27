package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"GoBlogClean/pkg/auth"
	"GoBlogClean/pkg/auth/input"
)

type UserHandler struct {
	userUsecase auth.UserUsecase
}

func NewUserHandler(userUsecase auth.UserUsecase) UserHandler {
	return UserHandler{userUsecase: userUsecase}
}

func (uh *UserHandler) Signup(c *gin.Context) {
	var requestBody *input.SignupRequest
	if err := c.BindJSON(&requestBody); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	if err := uh.userUsecase.Signup(requestBody); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": " succeeded"})
}

func (uh *UserHandler) Login(c *gin.Context) {
	var requestBody *input.LoginRequest
	if err := c.BindJSON(&requestBody); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	loginResponse, err := uh.userUsecase.Login(requestBody)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, loginResponse)
}

func (uh *UserHandler) GetUsers(c *gin.Context) {
	users, err := uh.userUsecase.GetUsers()
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, users)
}

func (uh *UserHandler) UpdateUsername(c *gin.Context) {
	var requestBody *input.UpdateUsernameRequest
	if err := c.BindJSON(&requestBody); err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	if err := uh.userUsecase.UpdateUsername(requestBody); err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "succeeded"})
}
