package groups

import (
	"encoding/json"
	"errors"
	"fmt"
	"mellow/models"
	"mellow/services"
	"mellow/utils"
	"net/http"
	"strings"
	"time"

	"github.com/google/uuid"
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

func InviteUserToGroup(groupSvc services.GroupService, notifSvc services.NotificationService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			utils.RespondError(w, http.StatusMethodNotAllowed, "Method not allowed", nil)
			return
		}

		id := strings.TrimPrefix(r.URL.Path, "/groups/invite/")

		if id == "" || strings.Contains(id, "/") {
			utils.RespondError(w, http.StatusNotFound, "Group not found", utils.ErrGroupNotFound)
			return
		}

		var payload struct {
			UserID string `json:"user_id"`
		}
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&payload); err != nil {
			utils.RespondError(w, http.StatusBadRequest, "Invalid payload", utils.ErrInvalidPayload)
			return
		}

		if payload.UserID == "" {
			utils.RespondError(w, http.StatusBadRequest, "User ID is required", utils.ErrInvalidPayload)
			return
		}

		userID, err := utils.GetUserIDFromContext(r.Context())
		if err != nil {
			utils.RespondError(w, http.StatusUnauthorized, "Unauthorized", utils.ErrUnauthorized)
			return
		}

		requestID, err := groupSvc.InviteUser(r.Context(), id, userID.String(), payload.UserID)
		if err != nil {
			switch {
			case errors.Is(err, utils.ErrGroupNotFound):
				utils.RespondError(w, http.StatusNotFound, "Group not found", utils.ErrGroupNotFound)
			case errors.Is(err, utils.ErrForbidden):
				utils.RespondError(w, http.StatusForbidden, "Forbidden", utils.ErrForbidden)
			case errors.Is(err, utils.ErrInvalidPayload):
				utils.RespondError(w, http.StatusBadRequest, "Invalid payload", utils.ErrInvalidPayload)
			default:
				utils.RespondError(w, http.StatusInternalServerError, "Failed to invite user", utils.ErrInternalServerError)
			}
			return
		}
		//Creer une notification pour l'invitation
		notif := &models.Notification{
			UserID:       uuid.MustParse(payload.UserID),
			SenderID:     userID.String(),
			RequestID:    &requestID,
			Type:         models.NotificationTypeGroupInvite,
			Seen:         false,
			CreationDate: time.Now(),
		}
		notifSvc.CreateNotification(r.Context(), notif)
		utils.RespondJSON(w, http.StatusOK, "User invited successfully", nil)
	}
}

func AnswerGroupInvite(groupSvc services.GroupService, userService services.UserService, notifSvc services.NotificationService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			utils.RespondError(w, http.StatusMethodNotAllowed, "Method not allowed", nil)
			return
		}

		requestId := strings.TrimPrefix(r.URL.Path, "/groups/invite/answer/")
		if requestId == "" {
			utils.RespondError(w, http.StatusBadRequest, "Request ID is required", utils.ErrInvalidPayload)
			return
		}

		action := r.URL.Query().Get("action")
		if action != "accept" && action != "reject" {
			utils.RespondError(w, http.StatusBadRequest, "Invalid action", utils.ErrBadRequest)
			return
		}

		request, err := userService.GetFollowRequestByID(r.Context(), requestId)
		if err != nil {
			utils.RespondError(w, http.StatusInternalServerError, "Failed to get follow request details", utils.ErrInternalServerError)
			return
		}
		userID, err := utils.GetUserIDFromContext(r.Context())
		if err != nil {
			utils.RespondError(w, http.StatusUnauthorized, "Unauthorized", utils.ErrUnauthorized)
			return
		}
		user, err := userService.GetUserByID(r.Context(), userID.String())
		if err != nil {
			utils.RespondError(w, http.StatusInternalServerError, "Failed to get user details", utils.ErrInternalServerError)
			return
		}

		if err := groupSvc.AnswerGroupInvite(r.Context(), *request, userID.String(), action); err != nil {
			fmt.Println("Error answering group invite:", err)
			switch {
			case errors.Is(err, utils.ErrInvalidPayload):
				utils.RespondError(w, http.StatusBadRequest, "Invalid payload", utils.ErrInvalidPayload)
			case errors.Is(err, utils.ErrForbidden):
				utils.RespondError(w, http.StatusForbidden, "Forbidden", utils.ErrForbidden)
			default:
				utils.RespondError(w, http.StatusInternalServerError, "Failed to answer group invite"+err.Error(), utils.ErrInternalServerError)
			}
			return
		}
		notif := &models.Notification{
			NotificationID:  uuid.New(),
			UserID:          request.SenderID,
			Seen:            false,
			CreationDate:    time.Now(),
			SenderID:        userID.String(),
			SenderUsername:  &user.Username,
			RequestID:       &request.RequestID,
			SenderAvatarURL: utils.GetFullImageURLAvatar(user.ImageURL),
		}
		switch action {
		case "accept":
			notif.Type = models.NotificationTypeAcceptedGroupInvite
		case "reject":
			notif.Type = models.NotificationTypeRejectedGroupInvite
		default:
			utils.RespondError(w, http.StatusBadRequest, "Invalid action", utils.ErrInvalidUserData)
			return
		}
		if err := notifSvc.CreateNotification(r.Context(), notif); err != nil {
			utils.RespondError(w, http.StatusInternalServerError, "Failed to create notification "+err.Error(), utils.ErrInternalServerError)
			return
		}

		utils.RespondJSON(w, http.StatusOK, "Group invite "+action+"ed successfully", nil)
	}
}
