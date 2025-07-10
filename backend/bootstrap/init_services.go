package bootstrap

import (
	"mellow/services"
	"mellow/services/servimpl"
)

type Services struct {
	UserService services.UserService
	AuthService services.AuthService
}

func InitServices(repos *Repositories) *Services {
	userService := servimpl.NewUserService(repos.UserRepository)
	authService := servimpl.NewAuthService(repos.UserRepository, repos.AuthRepository, userService)
	return &Services{
		UserService: userService,
		AuthService: authService,
	}
}
