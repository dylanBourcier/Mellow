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
	if user.Email == "" || user.Username == "" || user.Password == "" || user.Firstname == "" || user.Lastname == "" || user.Birthdate.IsZero() || user.Privacy == "" {
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

func (s *userServiceImpl) SendFollowRequest(ctx context.Context, senderID, receiverID string) (uuid.UUID, error) {
	if senderID == "" || receiverID == "" || senderID == receiverID {
		return uuid.Nil, fmt.Errorf("%s: invalid follow request", utils.ErrInvalidUserData)
	}
	var followRequest models.FollowRequest
	followRequest.SenderID = uuid.MustParse(senderID)
	followRequest.ReceiverID = uuid.MustParse(receiverID)
	if followRequest.SenderID == uuid.Nil || followRequest.ReceiverID == uuid.Nil {
		return uuid.Nil, fmt.Errorf("%s: invalid user IDs", utils.ErrInvalidUserData)
	}
	// Vérifier si la demande de suivi existe déjà
	exists, err := s.userRepo.IsFollowing(ctx, senderID, receiverID)
	if err != nil {
		return uuid.Nil, fmt.Errorf("failed to check if follow request exists: %w", err)
	}
	if exists {
		return uuid.Nil, fmt.Errorf("%s: follow request already exists", utils.ErrFollowRequestExists)
	}
	// Genereration de l'ID de la demande de suivi
	followRequest.RequestID = uuid.New()
	followRequest.CreationDate = time.Now() // Par défaut, la demande est en attente

	if err := s.userRepo.SendFollowRequest(ctx, followRequest); err != nil {
		return uuid.Nil, fmt.Errorf("failed to send follow request: %w", err)
	}
	fmt.Println("Follow request sent successfully: (service)", followRequest.RequestID)
	return followRequest.RequestID, nil
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

func (s *userServiceImpl) GetFollowers(ctx context.Context, viewerID, userID string) ([]*models.UserProfileData, error) {
	if userID == "" {
		return nil, fmt.Errorf("%s: empty id", utils.ErrUserNotFound)
	}
	users, err := s.userRepo.GetFollowers(ctx, viewerID, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve followers: %w", err)
	}
	return users, nil
}

func (s *userServiceImpl) GetFollowing(ctx context.Context, viewerID, userID string) ([]*models.UserProfileData, error) {
	if userID == "" {
		return nil, fmt.Errorf("%s: empty id", utils.ErrUserNotFound)
	}
	users, err := s.userRepo.GetFollowing(ctx, viewerID, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve following: %w", err)
	}
	return users, nil
}

func (s *userServiceImpl) SearchUsers(ctx context.Context, viewerID string, query string, groupId string, excludeGroupMembers bool) ([]*models.UserProfileData, error) {
	if query == "" {
		return []*models.UserProfileData{}, nil
	}
	if len(query) < 2 {
		return nil, fmt.Errorf("%s: query must be at least 3 characters long", utils.ErrInvalidUserData)
	}
	if groupId != "" && excludeGroupMembers {
		// If groupId is provided, we need to find users not in the group
		users, err := s.userRepo.SearchUsersExcludingGroupMembers(ctx, sanitize.SanitizeInput(query), groupId)
		if err != nil {
			return nil, fmt.Errorf("failed to search users excluding group members: %w", err)
		}
		return users, nil
	}
	//getUserID

	//If groupId and not excluding group members, search only the group members
	if groupId != "" && !excludeGroupMembers {
		// Do we need to search users in the group?
		// This part is commented out because it is not implemented yet.
		// users, err := s.userRepo.SearchUsersInGroup(ctx, sanitize.SanitizeInput(query), groupId)
		// if err != nil {
		// 	return nil, fmt.Errorf("failed to search users in group: %w", err)
		// }
		// return users, nil

	}
	users, err := s.userRepo.SearchUsers(ctx, viewerID, sanitize.SanitizeInput(query))
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

func (s *userServiceImpl) InsertFollow(ctx context.Context, followerID, followedID string) error {
	if followerID == "" || followedID == "" || followerID == followedID {
		return fmt.Errorf("%s: invalid follow", utils.ErrInvalidUserData)
	}
	if err := s.userRepo.InsertFollow(ctx, followerID, followedID); err != nil {
		return fmt.Errorf("failed to insert follow: %w", err)
	}
	return nil
}

func (s *userServiceImpl) GetUserPrivacy(ctx context.Context, userID string) (string, error) {
	if userID == "" {
		return "", fmt.Errorf("%s: empty id", utils.ErrInvalidUserData)
	}
	privacy, err := s.userRepo.GetUserPrivacy(ctx, userID)
	if err != nil {
		return "", fmt.Errorf("failed to retrieve user privacy settings: %w", err)
	}
	if privacy == "" {
		return "", utils.ErrUserNotFound
	}
	if privacy != "public" && privacy != "private" {
		return "", fmt.Errorf("%s: invalid privacy setting", utils.ErrInvalidUserData)
	}
	return privacy, nil
}

func (s *userServiceImpl) GetFollowRequestByID(ctx context.Context, requestID string) (*models.FollowRequest, error) {
	if requestID == "" {
		return nil, fmt.Errorf("%s: empty request ID", utils.ErrInvalidUserData)
	}
	request, err := s.userRepo.GetFollowRequestByID(ctx, requestID)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve follow request: %w", err)
	}
	if request == nil {
		return nil, utils.ErrFollowRequestNotFound
	}
	return request, nil
}

func (s *userServiceImpl) AnswerFollowRequest(ctx context.Context, request models.FollowRequest, userId, action string) error {
	if request.RequestID == uuid.Nil || userId == "" {
		return fmt.Errorf("%s: invalid request or user ID", utils.ErrInvalidUserData)
	}
	if request.ReceiverID != uuid.MustParse(userId) {
		return fmt.Errorf("%s: user not authorized to answer this request", utils.ErrUnauthorized)
	}

	if action != "accept" && action != "reject" {
		return fmt.Errorf("%s: invalid action", utils.ErrInvalidUserData)
	}

	if err := s.userRepo.AnswerFollowRequest(ctx, request, userId, action); err != nil {
		return fmt.Errorf("failed to answer follow request: %w", err)
	}

	return nil
}

func (s *userServiceImpl) IsFollowRequestExists(ctx context.Context, senderID, targetID string) (bool, error) {
	if senderID == "" || targetID == "" || senderID == targetID {
		return false, fmt.Errorf("%s: invalid follow request", utils.ErrInvalidUserData)
	}
	exists, err := s.userRepo.IsFollowRequestExists(ctx, senderID, targetID)
	if err != nil {
		return false, fmt.Errorf("failed to check if follow request exists: %w", err)
	}
	// Also check if the follow relationship already exists
	isFollowing, err := s.userRepo.IsFollowing(ctx, senderID, targetID)
	if err != nil {
		return false, fmt.Errorf("failed to check if user is following: %w", err)
	}
	if isFollowing || exists {
		return true, nil // If the user is already following or follow request already exists
	}
	return false, nil // No follow request or follow relationship exists

}
