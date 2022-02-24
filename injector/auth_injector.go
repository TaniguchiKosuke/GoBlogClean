package injector

import (
	"GoBlogClean/auth"
	"GoBlogClean/auth/handler"
	"GoBlogClean/auth/repository"
	"GoBlogClean/auth/usecase"
)

func InjectUserRepository() auth.UserRepository {
	return repository.NewUserRepository(InjectDB())
}

func InjectUserUsecase() auth.UserUsecase {
	return usecase.NewUserUsecase(InjectUserRepository())
}

func InjectUserHandler() handler.UserHandler {
	return handler.NewUserHandler(InjectUserUsecase())
}
