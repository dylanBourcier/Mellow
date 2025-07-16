package groups

import (
	"mellow/models"
	"mellow/services"
	"mellow/utils"
	"net/http"
	"strings"

	"github.com/google/uuid"
)

func GetAllGroups(w http.ResponseWriter, r *http.Request) {
	// TODO: liste des groupes
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
		//TODO: remove fake data
		var groupFake1 models.Group
		groupFake1.GroupID = uuid.New()
		groupFake1.Title = "Group Fake 1"
		var groupFake2 models.Group
		groupFake2.GroupID = uuid.New()
		groupFake2.Title = "Group Fake 2"
		groups = append(groups, &groupFake1, &groupFake2)

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
