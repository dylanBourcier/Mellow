package auth

import (
	"encoding/json"
	"errors"
	"mellow/config"
	"mellow/models"
	"mellow/services"
	"mellow/utils"
	"net/http"
	"time"
)

func SignUpHandler(userService services.UserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		if r.Method != http.MethodPost {
			utils.RespondError(w, http.StatusMethodNotAllowed, "Method not allowed", utils.ErrMethodNotAllowed)
			return
		}
		var user models.User
		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			utils.RespondError(w, http.StatusBadRequest, "Invalid JSON format: "+err.Error(), utils.ErrInvalidJSON)
			return
		}
		if err := userService.CreateUser(r.Context(), &user); err != nil {
			utils.RespondError(w, http.StatusBadRequest, "Invalid JSON payload: "+err.Error(), utils.ErrInvalidPayload)
			return
		}
		utils.RespondJSON(w, http.StatusCreated, "User created successfully", nil)
	}
}

func LoginHandler(authSvc services.AuthService) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			utils.RespondError(w, http.StatusMethodNotAllowed, "Method not allowed", utils.ErrMethodNotAllowed)
			return
		}

		var p models.LoginRequest
		if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
			utils.RespondError(w, http.StatusBadRequest, "Invalid JSON", utils.ErrInvalidJSON)
			return
		}

		user, sid, err := authSvc.Login(r.Context(), p.Identifier, p.Password)
		if err != nil {
			if errors.Is(err, utils.ErrUserNotFound) || errors.Is(err, utils.ErrInvalidCredentials) {
				// Si l'utilisateur n'existe pas ou si les identifiants sont incorrects
				// On ne donne pas d'indice sur l'existence de l'utilisateur pour éviter les attaques de type enumeration
				utils.RespondError(w, http.StatusUnauthorized, "Bad credentials", utils.ErrInvalidCredentials)
			} else {
				utils.RespondError(w, http.StatusInternalServerError, "Internal error", utils.ErrInternalServerError)
			}
			return
		}

		// Secure cookie
		http.SetCookie(w, &http.Cookie{
			Name:     config.CookieName,
			Value:    sid,
			Path:     "/",
			Expires:  time.Now().Add(config.CookieLifetime), // 7 days by default
			HttpOnly: true,
			Secure:   config.CookieSecure, // true in production (HTTPS)
			SameSite: http.SameSiteLaxMode,
		})

		utils.RespondJSON(w, http.StatusOK, "Login successful", map[string]any{"user_id": user.UserID, "session_id": sid})
	}
}

func LogoutHandler(authSvc services.AuthService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Verify if the method is POST
		if r.Method != http.MethodPost {
			utils.RespondError(w, http.StatusMethodNotAllowed, "Method not allowed", utils.ErrMethodNotAllowed)
			return
		}
		cookie, err := r.Cookie(config.CookieName)
		if err == nil {
			// Si le cookie existe, on tente de le supprimer
			http.SetCookie(w, &http.Cookie{Name: config.CookieName, Value: "", Path: "/", Expires: time.Unix(0, 0)})

			err := authSvc.Logout(r.Context(), cookie.Value)
			if err != nil {
				utils.RespondError(w, http.StatusInternalServerError, "Failed to log out: "+err.Error(), utils.ErrInternalServerError)
				return
			}

		}
		utils.RespondJSON(w, http.StatusOK, "Logged out", nil)
	}
}
