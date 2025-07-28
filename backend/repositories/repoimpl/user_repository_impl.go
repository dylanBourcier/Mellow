package repoimpl

import (
	"context"
	"database/sql"
	"fmt"
	"mellow/models"
	"mellow/repositories"
	"time"
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
	query := `SELECT user_id, email, password, username, firstname, lastname, birthdate, role, image_url, creation_date, description FROM users WHERE username = ?`
	var user models.User
	err := r.db.QueryRowContext(ctx, query, username).Scan(
		&user.UserID, &user.Email, &user.Password, &user.Username, &user.Firstname, &user.Lastname, &user.Birthdate, &user.Role, &user.ImageURL, &user.CreationDate, &user.Description)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("error retrieving user: %w", err)
	}
	return &user, nil
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
	query := `UPDATE users SET email = ?, password = ?, username = ?, firstname = ?, lastname = ?, birthdate = ?, role = ?, image_url = ?, description = ? WHERE user_id = ?`
	_, err := r.db.ExecContext(ctx, query,
		user.Email, user.Password, user.Username, user.Firstname, user.Lastname, user.Birthdate, user.Role, user.ImageURL, user.Description, user.UserID)
	if err != nil {
		return fmt.Errorf("error updating user: %w", err)
	}
	return nil
}

func (r *userRepositoryImpl) DeleteUser(ctx context.Context, userID string) error {
	_, err := r.db.ExecContext(ctx, `DELETE FROM users WHERE user_id = ?`, userID)
	if err != nil {
		return fmt.Errorf("error deleting user: %w", err)
	}
	return nil
}

func (r *userRepositoryImpl) InsertFollow(ctx context.Context, follower_id, followed_id string) error {
	query := `INSERT INTO follows (follower_id, followed_id, creation_date) VALUES (?, ?, ?)`
	_, err := r.db.ExecContext(ctx, query, follower_id, followed_id, time.Now())
	if err != nil {
		return fmt.Errorf("error inserting follow: %w", err)
	}
	return nil
}

func (r *userRepositoryImpl) SendFollowRequest(ctx context.Context, followRequest models.FollowRequest) error {
	query := `INSERT INTO follow_requests (request_id,sender_id,receiver_id, status, creation_date, type) VALUES (?,?, ?, 1, CURRENT_TIMESTAMP, 'user')`
	_, err := r.db.ExecContext(ctx, query, followRequest.RequestID, followRequest.SenderID, followRequest.ReceiverID)
	if err != nil {
		return fmt.Errorf("error following user: %w", err)
	}
	return nil
}

func (r *userRepositoryImpl) Unfollow(ctx context.Context, followerID, targetID string) error {
	// TODO: DELETE FROM follow WHERE follower_id = ? AND target_id = ?
	query := `DELETE FROM follow_requests WHERE sender_id = ? AND receiver_id = ?`
	_, err := r.db.ExecContext(ctx, query, followerID, targetID)
	if err != nil {
		return fmt.Errorf("error unfollowing user: %w", err)
	}
	return nil
}

func (r *userRepositoryImpl) GetFollowers(ctx context.Context, userID string) ([]*models.User, error) {
	query := `select u.user_id, u.email, u.password, u.username, u.firstname, u.lastname, u.birthdate, u.role, u.image_url, u.creation_date, u.description
			from users u
			join follow_requests f on f.sender_id = u.user_id
			where f.receiver_id = ? and f.type = 'user' and f.status = 1`
	rows, err := r.db.QueryContext(ctx, query, userID)
	if err != nil {
		return nil, fmt.Errorf("error retrieving followers: %w", err)
	}
	defer rows.Close()

	var users []*models.User
	for rows.Next() {
		var u models.User
		if err := rows.Scan(&u.UserID, &u.Email, &u.Password, &u.Username, &u.Firstname, &u.Lastname, &u.Birthdate, &u.Role, &u.ImageURL, &u.CreationDate, &u.Description); err != nil {
			return nil, fmt.Errorf("error scanning follower: %w", err)
		}
		users = append(users, &u)
	}
	return users, nil
}

func (r *userRepositoryImpl) GetFollowing(ctx context.Context, userID string) ([]*models.User, error) {
	query := `SELECT u.user_id, u.email, u.password, u.username, u.firstname, u.lastname, u.birthdate, u.role, u.image_url, u.creation_date, u.description
                  FROM users u
                  JOIN follow_requests f ON f.receiver_id = u.user_id
                  WHERE f.sender_id = ? AND f.type = 'user' AND f.status = 1`
	rows, err := r.db.QueryContext(ctx, query, userID)
	if err != nil {
		return nil, fmt.Errorf("error retrieving following: %w", err)
	}
	defer rows.Close()

	var users []*models.User
	for rows.Next() {
		var u models.User
		if err := rows.Scan(&u.UserID, &u.Email, &u.Password, &u.Username, &u.Firstname, &u.Lastname, &u.Birthdate, &u.Role, &u.ImageURL, &u.CreationDate, &u.Description); err != nil {
			return nil, fmt.Errorf("error scanning following: %w", err)
		}
		users = append(users, &u)
	}
	return users, nil
}

func (r *userRepositoryImpl) SearchUsers(ctx context.Context, query string) ([]*models.User, error) {
	like := "%" + query + "%"
	rows, err := r.db.QueryContext(ctx, `SELECT user_id, email, password, username, firstname, lastname, birthdate, role, image_url, creation_date, description 
												FROM users WHERE username LIKE ? OR firstname LIKE ? OR lastname LIKE ?`, like, like, like)
	if err != nil {
		return nil, fmt.Errorf("error searching users: %w", err)
	}
	defer rows.Close()

	var users []*models.User
	for rows.Next() {
		var u models.User
		if err := rows.Scan(&u.UserID, &u.Email, &u.Password, &u.Username, &u.Firstname, &u.Lastname, &u.Birthdate, &u.Role, &u.ImageURL, &u.CreationDate, &u.Description); err != nil {
			return nil, fmt.Errorf("error scanning user: %w", err)
		}
		users = append(users, &u)
	}
	return users, nil
}

func (r *userRepositoryImpl) IsFollowing(ctx context.Context, followerID, targetID string) (bool, error) {
	var exists bool
	query := `SELECT EXISTS(SELECT 1 FROM follows WHERE follower_id = ? AND followed_id = ?)`
	err := r.db.QueryRowContext(ctx, query, followerID, targetID).Scan(&exists)
	if err != nil {
		return false, fmt.Errorf("error checking following status: %w", err)
	}
	return exists, nil
}
func (r *userRepositoryImpl) GetUserProfile(ctx context.Context, viewerID, userID string) (*models.UserProfileData, error) {
	query := `SELECT u.user_id, u.username, u.firstname, u.lastname, u.email, u.birthdate, u.image_url, u.description, 
					 (SELECT COUNT(*) FROM follows f WHERE f.followed_id = u.user_id) AS followers_count,
					 (SELECT COUNT(*) FROM follows f WHERE f.follower_id = u.user_id) AS followed_count
			  FROM users u WHERE u.user_id = ?`
	var profile models.UserProfileData
	err := r.db.QueryRowContext(ctx, query, userID).Scan(
		&profile.UserID, &profile.Username, &profile.Firstname,
		&profile.Lastname, &profile.Email, &profile.Birthdate,
		&profile.ImageURL, &profile.Description, &profile.FollowersCount, &profile.FollowedCount)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // User not found
		}
		return nil, fmt.Errorf("error retrieving user profile: %w", err)
	}

	// Determine follow status
	followStatus, err := r.getFollowStatus(ctx, viewerID, userID)
	if err != nil {
		return nil, fmt.Errorf("error determining follow status: %w", err)
	}
	profile.FollowStatus = followStatus

	return &profile, nil
}

func (r *userRepositoryImpl) GetUserPrivacy(ctx context.Context, userID string) (string, error) {
	query := `SELECT privacy FROM users WHERE user_id = ?`
	var privacy string
	err := r.db.QueryRowContext(ctx, query, userID).Scan(&privacy)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", fmt.Errorf("user privacy not found")
		}
		return "", fmt.Errorf("error retrieving user privacy: %w", err)
	}
	return privacy, nil
}

func (r *userRepositoryImpl) getFollowStatus(ctx context.Context, viewerID, userID string) (string, error) {
	if viewerID == userID {
		return "yourself", nil
	}

	var followStatus string
	followQuery := `SELECT 'follows' AS status FROM follows WHERE follower_id = ? AND followed_id = ?
					UNION ALL
					SELECT 'requested' AS status FROM follow_requests WHERE sender_id = ? AND receiver_id = ?
					LIMIT 1`
	err := r.db.QueryRowContext(ctx, followQuery, viewerID, userID, viewerID, userID).Scan(&followStatus)
	if err != nil {
		if err == sql.ErrNoRows {
			return "not_follow", nil
		}
		return "", err
	}
	return followStatus, nil
}
