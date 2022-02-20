package handler

import (
	"GoBlogClean/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type UserHandler struct {
	userUsecase models.UserUsecase
}

func NewUserHandler(userUsecase models.UserUsecase) UserHandler {
	return UserHandler{userUsecase: userUsecase}
}

func (uh *UserHandler) Signup(c *gin.Context) {
	var user *models.User
	if err := c.BindJSON(&user); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	hashedBytePassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 12)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	user.Password = string(hashedBytePassword)

	if err := uh.userUsecase.SignUp(user); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "successed"})
}

func (uh *UserHandler) Login(c *gin.Context) {

}