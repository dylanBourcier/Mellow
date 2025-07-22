package groups

import (
	"mellow/services"
	"mellow/utils"
	"net/http"
	"strings"
)

func GetAllGroups(groupSvc services.GroupService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		groups, err := groupSvc.GetAllGroups(r.Context())
		if err != nil {
			utils.RespondError(w, http.StatusInternalServerError, "Failed to get groups: "+err.Error(), utils.ErrInternalServerError)
			return
		}

		if len(groups) == 0 {
			utils.RespondJSON(w, http.StatusOK, "No groups found", nil)
			return
		}

		utils.RespondJSON(w, http.StatusOK, "Groups retrieved successfully", groups)
	}
}
func GetAllGroupsWithoutUser(groupSvc services.GroupService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userID, err := utils.GetUserIDFromContext(r.Context())
		if err != nil {
			utils.RespondError(w, http.StatusUnauthorized, "Unauthorized", utils.ErrUnauthorized)
			return
		}

		groups, err := groupSvc.GetAllGroupsWithoutUser(r.Context(), userID.String())
		if err != nil {
			utils.RespondError(w, http.StatusInternalServerError, "Failed to get groups: "+err.Error(), utils.ErrInternalServerError)
			return
		}

		if len(groups) == 0 {
			utils.RespondJSON(w, http.StatusOK, "No groups found", nil)
			return
		}

		utils.RespondJSON(w, http.StatusOK, "Groups retrieved successfully", groups)
	}
}

func GetGroupsJoinedByUser(groupSvc services.GroupService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			utils.RespondError(w, http.StatusMethodNotAllowed, "Method not allowed", utils.ErrMethodNotAllowed)
			return
		}
		userID, err := utils.GetUserIDFromContext(r.Context())
		if err != nil {
			utils.RespondError(w, http.StatusUnauthorized, "Unauthorized", utils.ErrUnauthorized)
			return
		}

		groups, err := groupSvc.GetGroupsJoinedByUser(r.Context(), userID.String())
		if err != nil {
			utils.RespondError(w, http.StatusInternalServerError, "Failed to get groups: "+err.Error(), utils.ErrInternalServerError)
			return
		}

		if len(groups) == 0 {
			utils.RespondJSON(w, http.StatusOK, "No groups found", nil)
			return
		}

		utils.RespondJSON(w, http.StatusOK, "Groups retrieved successfully", groups)
	}

}

func GetGroupPosts(w http.ResponseWriter, r *http.Request, id string) {
	// TODO: posts du groupe
}

func GroupEventsHandler(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/groups/events/")
	if r.Method != http.MethodGet || id == "" || strings.Contains(id, "/") {
		http.NotFound(w, r)
		return
	}
	// TODO: événements du groupe
}

func GroupChatHandler(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/groups/chat/")
	if r.Method != http.MethodGet || id == "" || strings.Contains(id, "/") {
		http.NotFound(w, r)
		return
	}
	// TODO: chat du groupe, plutôt à faire dans la partie message ??
}
