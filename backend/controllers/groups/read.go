package groups

import (
	"fmt"
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

func GetGroupPosts(groupSvc services.GroupService, postSvc services.PostService, id string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userID, err := utils.GetUserIDFromContext(r.Context())
		if err != nil {
			utils.RespondError(w, http.StatusUnauthorized, "Unauthorized", utils.ErrUnauthorized)
			return
		}
		//Verifier que l'utilisateur a accès au groupe
		isMember, err := groupSvc.IsMember(r.Context(), id, userID.String())
		if err != nil {
			utils.RespondError(w, http.StatusInternalServerError, "Failed to check group membership: "+err.Error(), utils.ErrInternalServerError)
			return
		}
		if !isMember {
			utils.RespondJSON(w, http.StatusOK, "Not member", nil)
			return
		}
		limit := 10 // Default limit
		offset := 0 // Default offset
		query := r.URL.Query()
		if l := query.Get("limit"); l != "" {
			fmt.Sscanf(l, "%d", &limit)
		}
		if o := query.Get("offset"); o != "" {
			fmt.Sscanf(o, "%d", &offset)
		}
		if limit <= 0 || offset < 0 {
			utils.RespondError(w, http.StatusBadRequest, "Invalid limit or offset", utils.ErrInvalidPayload)
			return
		}
		_, err = groupSvc.GetGroupByID(r.Context(), id)
		if err != nil {
			utils.RespondError(w, http.StatusNotFound, "Group not found", utils.ErrGroupNotFound)
			return
		}

		posts, err := postSvc.GetGroupPosts(r.Context(), id, limit, offset)
		if err != nil {
			utils.RespondError(w, http.StatusInternalServerError, "Failed to get group posts: "+err.Error(), utils.ErrInternalServerError)
			return
		}
		utils.RespondJSON(w, http.StatusOK, "Group posts retrieved successfully", posts)
		return
	}
}

// func GroupEventsHandler(w http.ResponseWriter, r *http.Request) {
// 	id := strings.TrimPrefix(r.URL.Path, "/groups/events/")
// 	if r.Method != http.MethodGet || id == "" || strings.Contains(id, "/") {
// 		http.NotFound(w, r)
// 		return
// 	}
// 	// TODO: événements du groupe
// }

func GroupChatHandler(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/groups/chat/")
	if r.Method != http.MethodGet || id == "" || strings.Contains(id, "/") {
		http.NotFound(w, r)
		return
	}
	// TODO: chat du groupe, plutôt à faire dans la partie message ??
}

func GetGroupByID(groupSvc services.GroupService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		groupID := strings.TrimPrefix(r.URL.Path, "/groups/")
		if groupID == "" || strings.Contains(groupID, "/") {
			utils.RespondError(w, http.StatusNotFound, "Group not found", utils.ErrGroupNotFound)
			return
		}

		group, err := groupSvc.GetGroupByID(r.Context(), groupID)
		if err != nil {
			utils.RespondError(w, http.StatusInternalServerError, "Failed to get group: "+err.Error(), utils.ErrInternalServerError)
			return
		}
		// Check if the user is a member of the group
		userID, err := utils.GetUserIDFromContext(r.Context())
		if err != nil {
			utils.RespondError(w, http.StatusUnauthorized, "Unauthorized", utils.ErrUnauthorized)
			return
		}
		isMember, err := groupSvc.IsMember(r.Context(), groupID, userID.String())
		if err != nil {
			utils.RespondError(w, http.StatusInternalServerError, "Failed to check group membership: "+err.Error(), utils.ErrInternalServerError)
			return
		}

		response := map[string]interface{}{
			"group":     group,
			"is_member": isMember,
		}
		utils.RespondJSON(w, http.StatusOK, "Group retrieved successfully", response)
	}
}
