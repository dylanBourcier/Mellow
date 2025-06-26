package admin

import (
	"net/http"
	"strings"
)

func GetAllReportsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Méthode non autorisée", http.StatusMethodNotAllowed)
		return
	}
	// TODO: retourner la liste des signalements
}

func ModerateReportHandler(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/admin/reports/moderate/")
	if r.Method != http.MethodPost || id == "" || strings.Contains(id, "/") {
		http.NotFound(w, r)
		return
	}
	// TODO: appliquer modération (suppression, etc.)
}

func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/admin/users/")
	if r.Method != http.MethodDelete || id == "" || strings.Contains(id, "/") {
		http.NotFound(w, r)
		return
	}
	// TODO: supprimer l’utilisateur avec l’id
}
