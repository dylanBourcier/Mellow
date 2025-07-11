package repoimpl

import (
	"context"
	"database/sql"
	"fmt"
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
	query := "INSERT INTO users (user_id,email,password,username,firstname,lastname,birthdate,role,image_url,creation_date,description) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
	_, err := r.db.ExecContext(ctx, query, user.UserID, user.Email, user.Password, user.Username, user.Firstname, user.Lastname, user.Birthdate, user.Role, user.ImageURL, user.CreationDate, user.Description)
	if err != nil {
		return fmt.Errorf("error inserting user: %w", err)
	}
	return nil
}

func (r *userRepositoryImpl) FindUserByID(ctx context.Context, userID string) (*models.User, error) {
	query := `SELECT user_id, email, password, username, firstname, lastname, birthdate, role, image_url, creation_date, description 
	          FROM users WHERE user_id = ?`
	var user models.User
	err := r.db.QueryRowContext(ctx, query, userID).Scan(
		&user.UserID, &user.Email, &user.Password, &user.Username,
		&user.Firstname, &user.Lastname, &user.Birthdate,
		&user.Role, &user.ImageURL, &user.CreationDate,
		&user.Description)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // User not found
		}
		return nil, fmt.Errorf("error retrieving user: %w", err)
	}
	return &user, nil
}

func (r *userRepositoryImpl) FindUserByUsername(ctx context.Context, username string) (*models.User, error) {
	// TODO: Requête SELECT pour récupérer un utilisateur par nom d'utilisateur
	return nil, nil
}
func (r *userRepositoryImpl) UserExistsByEmailOrUsername(ctx context.Context, email, username string) (bool, error) {
	var count int
	err := r.db.QueryRowContext(ctx,
		`SELECT COUNT(*) FROM users WHERE email = ? OR username = ?`, email, username).Scan(&count)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}
func (r *userRepositoryImpl) GetUserByUsernameOrEmail(ctx context.Context, usernameOrEmail string) (*models.User, error) {
	var user models.User
	query := `SELECT user_id, email, password, username, firstname, lastname, birthdate, role, image_url, creation_date, description 
	          FROM users WHERE username = ? OR email = ?`
	err := r.db.QueryRowContext(ctx, query, usernameOrEmail, usernameOrEmail).Scan(
		&user.UserID, &user.Email, &user.Password, &user.Username,
		&user.Firstname, &user.Lastname, &user.Birthdate,
		&user.Role, &user.ImageURL, &user.CreationDate,
		&user.Description)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // User not found
		}
		return nil, fmt.Errorf("error retrieving user: %w", err)
	}
	return &user, nil
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
