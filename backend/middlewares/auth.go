package middlewares

import (
	"context"
	"log"

	"mellow/config"
	"mellow/repositories"
	"mellow/utils"
	"net/http"

	"github.com/google/uuid"
)

func AuthMiddleware(sessionRepo repositories.AuthRepository) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			cookie, err := r.Cookie(config.CookieName)
			if err != nil {
				utils.RespondError(w, http.StatusUnauthorized, "Missing session", utils.ErrUnauthorized)
				return
			}
			sessionID := cookie.Value

			uid, err := sessionRepo.GetUserIDFromSession(r.Context(), sessionID)
			if err != nil {
				utils.RespondError(w, http.StatusUnauthorized, "Invalid session", utils.ErrUnauthorized)
				return
			}
			if uid == uuid.Nil {
				utils.RespondError(w, http.StatusUnauthorized, "Invalid session", utils.ErrUnauthorized)
				return
			}
			if err := sessionRepo.UpdateLastActivity(r.Context(), sessionID); err != nil {
				log.Printf("Failed to update session last_activity: %v", err)
			}
			// Injecte le userID dans le contexte pour les handlers suivants
			ctx := context.WithValue(r.Context(), utils.CtxKeyUserID, uid)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
