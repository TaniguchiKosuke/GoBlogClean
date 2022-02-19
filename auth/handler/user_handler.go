package handler

import "GoBlogClean/models"

type UserHandler struct {
	userUsecase models.UserUsecase
}

func NewUserHandler(userUsecase models.UserUsecase) UserHandler {
	return UserHandler{userUsecase: userUsecase}
}