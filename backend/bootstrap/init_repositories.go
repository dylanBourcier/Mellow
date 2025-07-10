package bootstrap

import (
	"database/sql"
	"mellow/repositories"
	"mellow/repositories/repoimpl"
)

type Repositories struct {
	UserRepository repositories.UserRepository
	AuthRepository repositories.AuthRepository
}

func InitRepositories(db *sql.DB) *Repositories {
	return &Repositories{
		UserRepository: repoimpl.NewUserRepository(db),
		AuthRepository: repoimpl.NewAuthRepository(db),
	}
}
