package groups

import (
	"net/http"
	"strings"
)

func JoinGroupHandler(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/groups/join/")
	if r.Method != http.MethodPost || id == "" || strings.Contains(id, "/") {
		http.NotFound(w, r)
		return
	}
	// TODO: rejoindre le groupe id
}

func LeaveGroupHandler(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/groups/leave/")
	if r.Method != http.MethodDelete || id == "" || strings.Contains(id, "/") {
		http.NotFound(w, r)
		return
	}
	// TODO: quitter le groupe id
}
