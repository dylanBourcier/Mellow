package users

import (
	"mellow/utils"
	"net/http"
)

func ReportHandler(w http.ResponseWriter, r *http.Request) {
	utils.RespondError(w, http.StatusNotImplemented, "Not implemented", utils.ErrNotImplemented)
}
