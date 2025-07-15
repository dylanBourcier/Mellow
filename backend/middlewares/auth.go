package middlewares

import (
	"context"
	"fmt"

	"mellow/config"
	"mellow/services"
	"mellow/utils"
	"net/http"
)

func AuthMiddleware(authService services.AuthService) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			cookie, err := r.Cookie(config.CookieName)
			if err != nil {
				utils.RespondError(w, http.StatusUnauthorized, "Missing session", utils.ErrUnauthorized)
				return
			}

			sessionID := cookie.Value
			fmt.Println("Session ID:", sessionID)

			user, err := authService.GetUserFromSession(r.Context(), sessionID)
			if err != nil {
				utils.RespondError(w, http.StatusUnauthorized, "Invalid session", utils.ErrUnauthorized)
				return
			}
			uid := user.UserID

			_ = authService.UpdateLastActivity(r.Context(), sessionID)

			ctx := context.WithValue(r.Context(), utils.CtxKeyUserID, uid)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
