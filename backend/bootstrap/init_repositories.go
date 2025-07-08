package bootstrap

import (
	"database/sql"
	"mellow/repositories"
	"mellow/repositories/repoimpl"
)

type Repositories struct {
	UserRepository repositories.UserRepository
}

func InitRepositories(db *sql.DB) *Repositories {
	return &Repositories{
		UserRepository: repoimpl.NewUserRepository(db),
	}
}
