package impl

import (
	"context"
	"database/sql"
	"mellow/models"
	"mellow/repositories"
)

type userRepositoryImpl struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) repositories.UserRepository {
	return &userRepositoryImpl{db: db}
}

func (r *userRepositoryImpl) InsertUser(ctx context.Context, user *models.User) error {
	// TODO: Insérer un nouvel utilisateur dans la table users
	return nil
}

func (r *userRepositoryImpl) FindUserByID(ctx context.Context, userID string) (*models.User, error) {
	// TODO: Requête SELECT pour récupérer un utilisateur par ID
	return nil, nil
}

func (r *userRepositoryImpl) FindUserByUsername(ctx context.Context, username string) (*models.User, error) {
	// TODO: Requête SELECT pour récupérer un utilisateur par nom d'utilisateur
	return nil, nil
}

func (r *userRepositoryImpl) UpdateUser(ctx context.Context, user *models.User) error {
	// TODO: Requête UPDATE pour mettre à jour les informations de l'utilisateur
	return nil
}

func (r *userRepositoryImpl) DeleteUser(ctx context.Context, userID string) error {
	// TODO: Requête DELETE pour supprimer un utilisateur
	return nil
}

func (r *userRepositoryImpl) Follow(ctx context.Context, followerID, targetID string) error {
	// TODO: INSERT INTO follow (follower_id, target_id) VALUES (?, ?)
	return nil
}

func (r *userRepositoryImpl) Unfollow(ctx context.Context, followerID, targetID string) error {
	// TODO: DELETE FROM follow WHERE follower_id = ? AND target_id = ?
	return nil
}

func (r *userRepositoryImpl) GetFollowers(ctx context.Context, userID string) ([]*models.User, error) {
	// TODO: Requête JOIN pour récupérer les followers de l'utilisateur
	return nil, nil
}

func (r *userRepositoryImpl) GetFollowing(ctx context.Context, userID string) ([]*models.User, error) {
	// TODO: Requête JOIN pour récupérer les utilisateurs suivis
	return nil, nil
}

func (r *userRepositoryImpl) SearchUsers(ctx context.Context, query string) ([]*models.User, error) {
	// TODO: SELECT * FROM users WHERE username LIKE ?
	return nil, nil
}
