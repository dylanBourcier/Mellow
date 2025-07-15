package servimpl

import (
	"context"
	"errors"
	"fmt"
	"mellow/models"
	"mellow/repositories"
	"mellow/services"
	"mellow/utils"
	"time"

	"github.com/google/uuid"
)

type authServiceImpl struct {
	userRepo    repositories.UserRepository
	authRepo    repositories.AuthRepository
	userService services.UserService
}

// NewAuthService crée une nouvelle instance de AuthService.
func NewAuthService(userRepo repositories.UserRepository, authRepo repositories.AuthRepository, userService services.UserService) services.AuthService {
	return &authServiceImpl{userRepo, authRepo, userService}
}

func (s *authServiceImpl) Login(ctx context.Context, emailOrUsername, password string) (*models.User, string, error) {

	if err := s.authRepo.DeleteExpiredSessions(ctx); err != nil {
		return nil, "", fmt.Errorf("failed to delete expired sessions: %w", err)
	}
	user, err := s.userService.GetUserByUsernameOrEmail(ctx, emailOrUsername)
	if err != nil {
		if errors.Is(err, utils.ErrUserNotFound) {
			return nil, "", utils.ErrUserNotFound
		}
		return nil, "", fmt.Errorf("failed to get user: %w", err)
	}
	if !utils.ComparePasswords(user.Password, password) {
		return nil, "", utils.ErrInvalidCredentials
	}

	sid := uuid.New()
	now := time.Now()
	sess := &models.Session{
		SessionID:    sid,
		UserID:       user.UserID,
		CreationDate: now,
		LastActivity: now,
	}
	if err := s.authRepo.CreateSession(ctx, sess); err != nil {
		return nil, "", fmt.Errorf("failed to create session: %w", err)
	}

	// TODO: Créer une session en base si les identifiants sont valides
	return user, sid.String(), nil
}

func (s *authServiceImpl) Logout(ctx context.Context, sessionID string) error {
	if err := s.authRepo.DeleteSession(ctx, sessionID); err != nil {
		return fmt.Errorf("failed to delete session: %w", err)
	}
	return nil
}

func (s *authServiceImpl) IsAuthenticated(ctx context.Context, sessionID string) (bool, error) {
	// TODO: Vérifier si la session existe et est valide (ex: non expirée)
	return false, nil
}

func (s *authServiceImpl) GetUserFromSession(ctx context.Context, sessionID string) (*models.User, error) {
	userId, err := s.authRepo.GetUserIDFromSession(ctx, sessionID)
	if err != nil {
		return nil, fmt.Errorf("failed to get session: %w", err)
	}
	if userId == uuid.Nil {
		return nil, utils.ErrUserNotFound
	}

	user, err := s.userRepo.FindUserByID(ctx, userId.String())
	if err != nil {
		return nil, fmt.Errorf("failed to get user from session: %w", err)
	}
	return user, nil
}

func (s *authServiceImpl) UpdateLastActivity(ctx context.Context, sessionID string) error {
	if err := s.authRepo.UpdateLastActivity(ctx, sessionID); err != nil {
		return fmt.Errorf("failed to update last activity: %w", err)
	}
	return nil
}
