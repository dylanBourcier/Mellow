package admin

import (
	"mellow/controllers/admin"
	"net/http"
)

func RegisterAdminRoutes(mux *http.ServeMux) {
	// Voir tous les reports
	mux.HandleFunc("/admin/reports", admin.GetAllReportsHandler)

	// Modérer un contenu
	mux.HandleFunc("/admin/reports/moderate/", admin.ModerateReportHandler)

	// Supprimer un utilisateur
	mux.HandleFunc("/admin/users/", admin.DeleteUserHandler)
}
