package injector

import (
	"GoBlogClean/pkg/auth"
	"GoBlogClean/pkg/auth/handler"
	"GoBlogClean/pkg/auth/repository"
	"GoBlogClean/pkg/auth/usecase"
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
