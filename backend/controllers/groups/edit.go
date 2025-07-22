package groups

import (
	"encoding/json"
	"errors"
	"mellow/models"
	"mellow/services"
	"mellow/utils"
	"net/http"
)

// UpdateGroup returns a handler that updates a group's information.
func UpdateGroup(groupSvc services.GroupService, id string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()

		var payload models.GroupEditPayload

		if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
			utils.RespondError(w, http.StatusBadRequest, "Invalid JSON", utils.ErrInvalidJSON)
			return
		}

		userID, err := utils.GetUserIDFromContext(r.Context())
		if err != nil {
			utils.RespondError(w, http.StatusUnauthorized, "Unauthorized", utils.ErrUnauthorized)
			return
		}

		err = groupSvc.UpdateGroup(r.Context(), id, userID.String(), payload.Title, payload.Description)
		if err != nil {
			switch {
			case errors.Is(err, utils.ErrGroupNotFound):
				utils.RespondError(w, http.StatusNotFound, "Group not found", err)
			case errors.Is(err, utils.ErrForbidden):
				utils.RespondError(w, http.StatusForbidden, "Forbidden", err)
			case errors.Is(err, utils.ErrContentTooLong), errors.Is(err, utils.ErrContentTooShort), errors.Is(err, utils.ErrInvalidPayload):
				utils.RespondError(w, http.StatusBadRequest, "Invalid group", err)
			default:
				utils.RespondError(w, http.StatusInternalServerError, "Failed to update group", err)
			}
			return
		}

		utils.RespondJSON(w, http.StatusOK, "Group updated successfully", nil)
	}
}
