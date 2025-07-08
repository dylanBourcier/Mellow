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
	return &Services{
		UserService: servimpl.NewUserService(repos.UserRepository),
		AuthService: servimpl.NewAuthService(repos.UserRepository, repos.AuthRepository),
	}
}
