package groups

import (
	"errors"
	"mellow/models"
	"mellow/services"
	"mellow/utils"
	"net/http"
)

func CreateGroup(groupSvc services.GroupService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseMultipartForm(10 << 20); err != nil { // 10MB max size
			utils.RespondError(w, http.StatusBadRequest, "Invalid multipart form data", utils.ErrInvalidPayload)
			return
		}

		userID, err := utils.GetUserIDFromContext(r.Context())
		if err != nil {
			utils.RespondError(w, http.StatusUnauthorized, "Unauthorized", utils.ErrUnauthorized)
			return
		}

		title := r.FormValue("title")
		if title == "" {
			utils.RespondError(w, http.StatusBadRequest, "Title is required", utils.ErrInvalidPayload)
			return
		}

		desc := r.FormValue("description")

		g := models.Group{
			Title:       title,
			Description: desc,
			UserID:      userID,
		}

		if err := groupSvc.CreateGroup(r.Context(), &g); err != nil {
			switch {
			case errors.Is(err, utils.ErrGroupAlreadyExists):
				utils.RespondError(w, http.StatusConflict, "Group already exists", utils.ErrGroupAlreadyExists)
			case errors.Is(err, utils.ErrInvalidPayload):
				utils.RespondError(w, http.StatusBadRequest, "Invalid data", utils.ErrInvalidPayload)
			default:
				utils.RespondError(w, http.StatusInternalServerError, "Failed to create group", utils.ErrInternalServerError)
			}
			return
		}

		utils.RespondJSON(w, http.StatusCreated, "Group created successfully", map[string]any{"group_id": g.GroupID})
	}
}
