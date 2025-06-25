package notifications

import "net/http"

func GetNotificationsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Méthode non autorisée", http.StatusMethodNotAllowed)
		return
	}
	// TODO: retourner la liste des notifications de l'utilisateur connecté
}
