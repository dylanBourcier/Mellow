package groups

import (
	"errors"
	"mellow/services"
	"mellow/utils"
	"net/http"
	"strings"
)

func JoinGroup(groupSvc services.GroupService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := strings.TrimPrefix(r.URL.Path, "/groups/join/")
		if id == "" || strings.Contains(id, "/") {
			utils.RespondError(w, http.StatusNotFound, "Groupe introuvable", utils.ErrGroupNotFound)
			return
		}

		userID, err := utils.GetUserIDFromContext(r.Context())
		if err != nil {
			utils.RespondError(w, http.StatusUnauthorized, "Unauthorized", utils.ErrUnauthorized)
			return
		}

		if _, err := groupSvc.GetGroupByID(r.Context(), id); err != nil {
			if errors.Is(err, utils.ErrGroupNotFound) {
				utils.RespondError(w, http.StatusNotFound, "Groupe introuvable", utils.ErrGroupNotFound)
			} else {
				utils.RespondError(w, http.StatusInternalServerError, "Erreur interne", utils.ErrInternalServerError)
			}
			return
		}

		if err := groupSvc.AddMember(r.Context(), id, userID.String()); err != nil {
			switch {
			case errors.Is(err, utils.ErrResourceConflict):
				utils.RespondError(w, http.StatusConflict, "Deja membre du groupe", utils.ErrResourceConflict)
			case errors.Is(err, utils.ErrInvalidPayload):
				utils.RespondError(w, http.StatusBadRequest, "Données invalides", utils.ErrInvalidPayload)
			default:
				utils.RespondError(w, http.StatusInternalServerError, "Impossible de rejoindre le groupe", utils.ErrInternalServerError)
			}
			return
		}

		utils.RespondJSON(w, http.StatusOK, "Groupe rejoint", nil)
	}
}

func LeaveGroupHandler(groupSvc services.GroupService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := strings.TrimPrefix(r.URL.Path, "/groups/leave/")
		if r.Method != http.MethodDelete || id == "" || strings.Contains(id, "/") {
			http.NotFound(w, r)
			return
		}

		userID, err := utils.GetUserIDFromContext(r.Context())
		if err != nil {
			utils.RespondError(w, http.StatusUnauthorized, "Unauthorized", utils.ErrUnauthorized)
			return
		}

		if err := groupSvc.RemoveMember(r.Context(), id, userID.String()); err != nil {
			switch err {
			case utils.ErrInvalidPayload:
				utils.RespondError(w, http.StatusBadRequest, "Invalid payload", err)
			case utils.ErrGroupNotFound:
				utils.RespondError(w, http.StatusNotFound, "Group not found", err)
			case utils.ErrForbidden:
				utils.RespondError(w, http.StatusForbidden, "Forbidden", err)
			default:
				utils.RespondError(w, http.StatusInternalServerError, "Failed to leave group", err)
			}
			return
		}

		utils.RespondJSON(w, http.StatusOK, "Left group successfully", nil)
	}
}
