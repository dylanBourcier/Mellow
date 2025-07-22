package groups

import (
	"errors"
	"mellow/services"
	"mellow/utils"
	"net/http"
)

func DeleteGroup(groupSvc services.GroupService, id string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			utils.RespondError(w, http.StatusMethodNotAllowed, "Method not allowed", utils.ErrMethodNotAllowed)
			return
		}
		userID, err := utils.GetUserIDFromContext(r.Context())
		if err != nil {
			utils.RespondError(w, http.StatusUnauthorized, "Unauthorized", utils.ErrUnauthorized)
			return
		}
		if err := groupSvc.DeleteGroup(r.Context(), id, userID.String()); err != nil {
			switch {
			case errors.Is(err, utils.ErrGroupNotFound):
				utils.RespondError(w, http.StatusNotFound, "Group not found", utils.ErrGroupNotFound)
			case errors.Is(err, utils.ErrUnauthorized):
				utils.RespondError(w, http.StatusForbidden, "Forbidden", utils.ErrUnauthorized)
			case errors.Is(err, utils.ErrResourceConflict):
				utils.RespondError(w, http.StatusBadRequest, "Group not empty", utils.ErrResourceConflict)
			case errors.Is(err, utils.ErrInvalidPayload):
				utils.RespondError(w, http.StatusBadRequest, "Invalid group ID", utils.ErrInvalidPayload)
			default:
				utils.RespondError(w, http.StatusInternalServerError, "Failed to delete group", err)
			}
			return
		}
		utils.RespondJSON(w, http.StatusOK, "Group deleted successfully", nil)
	}
}
