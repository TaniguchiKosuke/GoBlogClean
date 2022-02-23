package handler

import (
	"GoBlogClean/common"
	"GoBlogClean/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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

	u, err := uuid.NewRandom()
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	uuidStr := u.String()
	user.ID = uuidStr
	// よく考えたらJWTはDBに保存しないかも
	jwtToken, err := common.CreateJWTToken(user)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	user.JwtToken = jwtToken

	if err := uh.userUsecase.Signup(user); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"Token": " succeeded"})
}

func (uh *UserHandler) Login(c *gin.Context) {

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
