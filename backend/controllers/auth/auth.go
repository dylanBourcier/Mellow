package auth

import (
	"encoding/json"
	"mellow/models"
	"mellow/services"
	"mellow/utils"
	"net/http"
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

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Méthode non autorisée", http.StatusMethodNotAllowed)
		return
	}
	// TODO: vérifier les identifiants et créer la session / cookie
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Méthode non autorisée", http.StatusMethodNotAllowed)
		return
	}
	// TODO: invalider la session côté serveur + supprimer le cookie
}
