package repoimpl

import (
	"context"
	"database/sql"
	"mellow/models"
	"mellow/repositories"
)

type authRepositoryImpl struct {
	db *sql.DB
}

func NewAuthRepository(db *sql.DB) repositories.AuthRepository {
	return &authRepositoryImpl{db: db}
}

func (r *authRepositoryImpl) CreateSession(ctx context.Context, session *models.Session) error {
	// TODO: INSERT INTO sessions (id, user_id, created_at, expires_at) VALUES (?, ?, ?, ?)
	return nil
}

func (r *authRepositoryImpl) DeleteSession(ctx context.Context, sessionID string) error {
	// TODO: DELETE FROM sessions WHERE id = ?
	return nil
}

func (r *authRepositoryImpl) GetUserFromSession(ctx context.Context, sessionID string) (*models.User, error) {
	// TODO: JOIN entre sessions et users pour récupérer l'utilisateur à partir de la session
	return nil, nil
}

func (r *authRepositoryImpl) IsSessionValid(ctx context.Context, sessionID string) (bool, error) {
	// TODO: Vérifier si la session existe et si expires_at > NOW()
	return false, nil
}
