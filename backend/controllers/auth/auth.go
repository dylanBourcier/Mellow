package auth

import (
	"encoding/json"
	"errors"
	"io"
	"mellow/config"
	"mellow/models"
	"mellow/services"
	"mellow/utils"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/google/uuid"
)

func SignUpHandler(userService services.UserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()

		if r.Method != http.MethodPost {
			utils.RespondError(w, http.StatusMethodNotAllowed, "Method not allowed", utils.ErrMethodNotAllowed)
			return
		}

		// Limite taille max (par ex 10 Mo)
		if err := r.ParseMultipartForm(10 << 20); err != nil {
			utils.RespondError(w, http.StatusBadRequest, "Failed to parse multipart form: "+err.Error(), utils.ErrInvalidPayload)
			return
		}

		// Récupère les champs texte
		user := models.User{}

		user.Email = r.FormValue("email")
		user.Username = r.FormValue("username")
		user.Firstname = r.FormValue("firstname")
		user.Lastname = r.FormValue("lastname")
		user.Description = nil
		if desc := r.FormValue("description"); desc != "" {
			user.Description = &desc
		}

		// Parse birthdate (exemple format "2006-01-02")
		bdStr := r.FormValue("birthdate")
		if bdStr != "" {
			birthdate, err := time.Parse("2006-01-02", bdStr)
			if err != nil {
				utils.RespondError(w, http.StatusBadRequest, "Invalid birthdate format", utils.ErrInvalidPayload)
				return
			}
			user.Birthdate = birthdate
		}

		// Password (à gérer selon ta struct, mais attention il faut l’inclure dans la struct)
		password := r.FormValue("password")
		if password == "" {
			utils.RespondError(w, http.StatusBadRequest, "Password is required", utils.ErrInvalidPayload)
			return
		}
		user.Password = password

		// Traite le fichier avatar
		file, header, err := r.FormFile("avatar")
		if err == nil {
			defer file.Close()

			// Génère un nom unique pour éviter les collisions
			filename := uuid.New().String() + filepath.Ext(header.Filename)
			savePath := filepath.Join("uploads", "avatars", filename)

			// Crée le dossier si besoin
			if err := os.MkdirAll(filepath.Dir(savePath), os.ModePerm); err != nil {
				utils.RespondError(w, http.StatusInternalServerError, "Unable to create directory", utils.ErrInternalServerError)
				return
			}

			dst, err := os.Create(savePath)
			if err != nil {
				utils.RespondError(w, http.StatusInternalServerError, "Unable to save avatar", utils.ErrInternalServerError)
				return
			}
			defer dst.Close()

			if _, err := io.Copy(dst, file); err != nil {
				utils.RespondError(w, http.StatusInternalServerError, "Failed to save avatar", utils.ErrInternalServerError)
				return
			}

			user.ImageURL = &filename
		} else if err != http.ErrMissingFile {
			// Erreur autre que "fichier manquant"
			utils.RespondError(w, http.StatusBadRequest, "Error reading avatar file: "+err.Error(), utils.ErrInvalidPayload)
			return
		}

		// Appelle le service (il hash le mdp, valide, insert)
		if err := userService.CreateUser(r.Context(), &user); err != nil {
			utils.RespondError(w, http.StatusBadRequest, "User creation failed: "+err.Error(), utils.ErrInvalidPayload)
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
