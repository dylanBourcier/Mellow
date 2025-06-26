package notifications

import (
	"net/http"
	"strings"
)

func MarkAsReadHandler(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/notifications/read/")
	if r.Method != http.MethodPost || id == "" || strings.Contains(id, "/") {
		http.NotFound(w, r)
		return
	}
	// TODO: marquer la notification id comme lue
}
