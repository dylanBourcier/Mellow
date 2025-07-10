package repoimpl

import (
	"context"
	"database/sql"
	"mellow/models"
	"mellow/repositories"

	"github.com/google/uuid"
)

type authRepositoryImpl struct {
	db *sql.DB
}

func NewAuthRepository(db *sql.DB) repositories.AuthRepository {
	return &authRepositoryImpl{db: db}
}

func (r *authRepositoryImpl) CreateSession(ctx context.Context, s *models.Session) error {
	query := `INSERT INTO sessions (session_id, user_id, creation_date, last_activity)
         VALUES (?, ?, ?, ?)`
	_, err := r.db.ExecContext(ctx, query, s.SessionID, s.UserID, s.CreationDate, s.LastActivity)
	return err
}

func (r *authRepositoryImpl) DeleteSession(ctx context.Context, sid string) error {
	_, err := r.db.ExecContext(ctx, `DELETE FROM sessions WHERE session_id = ?`, sid)
	return err
}

func (r *authRepositoryImpl) GetUserFromSession(ctx context.Context, sessionID string) (*models.User, error) {
	// TODO: JOIN entre sessions et users pour récupérer l'utilisateur à partir de la session
	return nil, nil
}
func (r *authRepositoryImpl) GetUserIDFromSession(ctx context.Context, sessionID string) (uuid.UUID, error) {
	query := `SELECT user_id FROM sessions WHERE session_id = ?`
	var idStr string
	err := r.db.QueryRowContext(ctx, query, sessionID).Scan(&idStr)
	if err != nil {
		if err == sql.ErrNoRows {
			return uuid.Nil, nil // Session not found
		}
		return uuid.Nil, err // Other error
	}
	return uuid.Parse(idStr)
}

func (r *authRepositoryImpl) IsSessionValid(ctx context.Context, sessionID string) (bool, error) {
	// TODO: Vérifier si la session existe et si expires_at > NOW()
	return false, nil
}

func (r *authRepositoryImpl) DeleteExpiredSessions(ctx context.Context) error {
	_, err := r.db.ExecContext(ctx,
		`DELETE FROM sessions WHERE last_activity <= datetime('now', '-7 days')`)
	return err
}

func (r *authRepositoryImpl) UpdateLastActivity(ctx context.Context, sessionID string) error {
	_, err := r.db.ExecContext(ctx,
		`UPDATE sessions SET last_activity = CURRENT_TIMESTAMP WHERE session_id = ?`, sessionID)
	return err
}
