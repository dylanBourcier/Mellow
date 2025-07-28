package servimpl

import (
	"context"
	"fmt"
	"mellow/models"
	"mellow/repositories"
	"mellow/services"
	"mellow/utils"
	"mellow/utils/sanitize"
	"time"

	"github.com/google/uuid"
)

type userServiceImpl struct {
	userRepo repositories.UserRepository // Référence au repository utilisateur
}

// NewUserService crée une nouvelle instance de UserService.
func NewUserService(userRepo repositories.UserRepository) services.UserService {
	return &userServiceImpl{userRepo: userRepo}
}

func (s *userServiceImpl) CreateUser(ctx context.Context, user *models.User) error {
	// Validation simple
	if user.Email == "" || user.Username == "" || user.Password == "" || user.Firstname == "" || user.Lastname == "" || user.Birthdate.IsZero() {
		return fmt.Errorf("%s: missing required fields", utils.ErrInvalidUserData)
	}
	// Vérifier si l'utilisateur existe déjà par email ou nom d'utilisateur
	exists, err := s.userRepo.UserExistsByEmailOrUsername(ctx, user.Email, user.Username)
	if err != nil {
		return fmt.Errorf("failed to check if user exists: %w", err)
	}
	if exists {
		return utils.ErrUserAlreadyExists
	}

	uuid, err := uuid.NewRandom()
	if err != nil {
		return fmt.Errorf("failed to generate user ID: %w", err)
	}
	user.UserID = uuid

	user.CreationDate = time.Now() // Assigner la date de création actuelle

	if user.Role == "" {
		user.Role = "user" // Assigner un rôle par défaut si non spécifié
	}

	//Hash the password before storing it
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return fmt.Errorf("failed to hash password: %w", err)
	}
	user.Password = hashedPassword

	return s.userRepo.InsertUser(ctx, user)
}

func (s *userServiceImpl) GetUserByID(ctx context.Context, userID string) (*models.User, error) {
	if userID == "" {
		return nil, fmt.Errorf("%s: empty id", utils.ErrInvalidUserData)
	}

	user, err := s.userRepo.FindUserByID(ctx, sanitize.SearchQuery(userID))
	if err != nil {
		return nil, fmt.Errorf("failed to find user by id: %w", err)
	}

	if user == nil {
		return nil, utils.ErrUserNotFound
	}
	return user, nil
}

func (s *userServiceImpl) GetUserByUsername(ctx context.Context, username string) (*models.User, error) {
	// TODO: Appliquer les éventuelles règles métier (ex: autorisations)
	if username == "" {
		return nil, fmt.Errorf("%s: empty username", utils.ErrUserNotFound)
	}

	user, err := s.userRepo.FindUserByUsername(ctx, sanitize.SearchQuery(username))
	if err != nil {
		return nil, fmt.Errorf("failed to find user by username: %w", err)
	}

	if user == nil {
		return nil, utils.ErrUserNotFound
	}
	return user, nil
}
func (s *userServiceImpl) GetUserByUsernameOrEmail(ctx context.Context, identifier string) (*models.User, error) {

	var user *models.User
	user, err := s.userRepo.GetUserByUsernameOrEmail(ctx, identifier)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve user: %w", err)
	}
	if user == nil {
		return nil, utils.ErrUserNotFound
	}
	return user, nil
}

func (s *userServiceImpl) UpdateUser(ctx context.Context, user *models.User) error {
	if user == nil || user.UserID == uuid.Nil {
		return fmt.Errorf("%s: empty id", utils.ErrUserNotFound)
	}
	if err := s.userRepo.UpdateUser(ctx, user); err != nil {
		return fmt.Errorf("failed to update user: %w", err)
	}
	return nil
}

func (s *userServiceImpl) DeleteUser(ctx context.Context, userID string) error {
	if userID == "" {
		return fmt.Errorf("%s: empty id", utils.ErrInvalidUserData)
	}
	if err := s.userRepo.DeleteUser(ctx, userID); err != nil {
		return fmt.Errorf("failed to delete user: %w", err)
	}
	return nil
}

func (s *userServiceImpl) Authenticate(ctx context.Context, username, password string) (*models.User, error) {
	user, err := s.userRepo.FindUserByUsername(ctx, username)
	if err != nil {
		return nil, fmt.Errorf("failed to get user: %w", err)
	}
	if user == nil {
		return nil, utils.ErrUserNotFound
	}
	if !utils.ComparePasswords(user.Password, password) {
		return nil, utils.ErrInvalidCredentials
	}
	return user, nil
}

func (s *userServiceImpl) FollowUser(ctx context.Context, followerID, targetID string) error {
	if followerID == "" || targetID == "" || followerID == targetID {
		return fmt.Errorf("%s: invalid follow", utils.ErrInvalidUserData)
	}
	if err := s.userRepo.Follow(ctx, followerID, targetID); err != nil {
		return fmt.Errorf("failed to follow user: %w", err)
	}
	return nil
}

func (s *userServiceImpl) UnfollowUser(ctx context.Context, followerID, targetID string) error {
	if followerID == "" || targetID == "" || followerID == targetID {
		return fmt.Errorf("%s: invalid unfollow", utils.ErrInvalidUserData)
	}
	if err := s.userRepo.Unfollow(ctx, followerID, targetID); err != nil {
		return fmt.Errorf("failed to unfollow user: %w", err)
	}
	return nil
}

func (s *userServiceImpl) GetFollowers(ctx context.Context, userID string) ([]*models.User, error) {
	if userID == "" {
		return nil, fmt.Errorf("%s: empty id", utils.ErrUserNotFound)
	}
	users, err := s.userRepo.GetFollowers(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve followers: %w", err)
	}
	return users, nil
}

func (s *userServiceImpl) GetFollowing(ctx context.Context, userID string) ([]*models.User, error) {
	if userID == "" {
		return nil, fmt.Errorf("%s: empty id", utils.ErrUserNotFound)
	}
	users, err := s.userRepo.GetFollowing(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve following: %w", err)
	}
	return users, nil
}

func (s *userServiceImpl) SearchUsers(ctx context.Context, query string) ([]*models.User, error) {
	if query == "" {
		return []*models.User{}, nil
	}
	users, err := s.userRepo.SearchUsers(ctx, sanitize.SanitizeInput(query))
	if err != nil {
		return nil, fmt.Errorf("failed to search users: %w", err)
	}
	return users, nil
}

func (s *userServiceImpl) IsFollowing(ctx context.Context, followerID, targetID string) (bool, error) {
	return s.userRepo.IsFollowing(ctx, followerID, targetID)
}

func (s *userServiceImpl) GetUserProfileData(ctx context.Context, viewerID, userID string) (*models.UserProfileData, error) {
	if userID == "" {
		return nil, fmt.Errorf("%s: empty id", utils.ErrInvalidUserData)
	}
	userProfileData, err := s.userRepo.GetUserProfile(ctx, viewerID, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve user profile: %w", err)
	}
	if userProfileData == nil {
		return nil, utils.ErrUserNotFound
	}
	if userProfileData.ImageURL != nil {
		userProfileData.ImageURL = utils.GetFullImageURLAvatar(userProfileData.ImageURL)
	}
	return userProfileData, nil

}
