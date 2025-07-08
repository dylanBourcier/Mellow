package bootstrap

import (
	"mellow/services"
	"mellow/services/servimpl"
)

type Services struct {
	UserService services.UserService
}

func InitServices(repos *Repositories) *Services {
	return &Services{
		UserService: servimpl.NewUserService(repos.UserRepository),
	}
}
