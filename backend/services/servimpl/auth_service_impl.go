package servimpl

import (
	"context"
	"database/sql"
	"mellow/models"
	"mellow/services"
)

type authServiceImpl struct {
	db *sql.DB
}

// NewAuthService crée une nouvelle instance de AuthService.
func NewAuthService(db *sql.DB) services.AuthService {
	return &authServiceImpl{db: db}
}

func (s *authServiceImpl) Login(ctx context.Context, username, password string) (*models.User, error) {
	// TODO: Récupérer l'utilisateur depuis le repository
	// TODO: Utiliser utils.ComparePasswords pour vérifier le mot de passe
	// TODO: Créer une session en base si les identifiants sont valides
	return nil, nil
}

func (s *authServiceImpl) Logout(ctx context.Context, sessionID string) error {
	// TODO: Supprimer la session de la base de données
	return nil
}

func (s *authServiceImpl) IsAuthenticated(ctx context.Context, sessionID string) (bool, error) {
	// TODO: Vérifier si la session existe et est valide (ex: non expirée)
	return false, nil
}

func (s *authServiceImpl) GetUserFromSession(ctx context.Context, sessionID string) (*models.User, error) {
	// TODO: Récupérer l'utilisateur lié à une session donnée
	return nil, nil
}
