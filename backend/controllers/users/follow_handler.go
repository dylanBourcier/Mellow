package users

import (
	"mellow/models"
	"mellow/services"
	"mellow/utils"
	"net/http"
	"strings"
	"time"

	"github.com/google/uuid"
)

func SendFollowRequest(userService services.UserService, notificationService services.NotificationService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			utils.RespondError(w, http.StatusMethodNotAllowed, "Method not allowed", utils.ErrMethodNotAllowed)
			return
		}

		targetID := strings.TrimPrefix(r.URL.Path, "/users/follow/")
		senderID, err := utils.GetUserIDFromContext(r.Context())
		if err != nil {
			utils.RespondError(w, http.StatusUnauthorized, "Unauthorized", utils.ErrUnauthorized)
			return
		}
		//Check if the target user is in public or private to send a notification and validate the request
		privacy, err := userService.GetUserPrivacy(r.Context(), targetID)
		if err != nil {
			utils.RespondError(w, http.StatusInternalServerError, "Failed to get user privacy settings"+err.Error(), utils.ErrInternalServerError)
			return
		}

		var notif *models.Notification
		if privacy == "private" {

			requestID, err := userService.SendFollowRequest(r.Context(), senderID.String(), targetID)
			if err != nil {
				utils.RespondError(w, http.StatusInternalServerError, "Failed to send follow request"+err.Error(), utils.ErrInternalServerError)
				return
			}
			notif = &models.Notification{
				UserID:       uuid.MustParse(targetID),
				RequestID:    &requestID,
				SenderID:     senderID.String(),
				Type:         models.NotificationTypeFollowRequest,
				Seen:         false,
				CreationDate: time.Now(),
			}
		} else if privacy == "public" {
			if err := userService.InsertFollow(r.Context(), senderID.String(), targetID); err != nil {
				utils.RespondError(w, http.StatusInternalServerError, "Failed to follow user"+err.Error(), utils.ErrInternalServerError)
				return
			}

			notif = &models.Notification{
				UserID:       uuid.MustParse(targetID),
				SenderID:     senderID.String(),
				Type:         models.NotificationTypeNewFollower,
				Seen:         false,
				CreationDate: time.Now(),
			}
		}

		if err := notificationService.CreateNotification(r.Context(), notif); err != nil {
			utils.RespondError(w, http.StatusInternalServerError, "Failed to create notification "+err.Error(), utils.ErrInternalServerError)
			return
		}

		utils.RespondJSON(w, http.StatusOK, "Follow request sent", nil)
	}
}

// func FollowUser(userService services.UserService) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		if r.Method != http.MethodPost {
// 			utils.RespondError(w, http.StatusMethodNotAllowed, "Method not allowed", utils.ErrMethodNotAllowed)
// 			return
// 		}

// 		targetID := strings.TrimPrefix(r.URL.Path, "/users/follow/")
// 		followerID, err := utils.GetUserIDFromContext(r.Context())
// 		if err != nil {
// 			utils.RespondError(w, http.StatusUnauthorized, "Unauthorized", utils.ErrUnauthorized)
// 			return
// 		}

// 		if err := userService.FollowUser(r.Context(), followerID.String(), targetID); err != nil {
// 			utils.RespondError(w, http.StatusInternalServerError, "Failed to follow user", utils.ErrInternalServerError)
// 			return
// 		}

// 		utils.RespondJSON(w, http.StatusOK, "User followed", nil)
// 	}
// }

func UnfollowUser(userService services.UserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			utils.RespondError(w, http.StatusMethodNotAllowed, "Method not allowed", utils.ErrMethodNotAllowed)
			return
		}

		targetID := strings.TrimPrefix(r.URL.Path, "/users/follow/")
		followerID, err := utils.GetUserIDFromContext(r.Context())
		if err != nil {
			utils.RespondError(w, http.StatusUnauthorized, "Unauthorized", utils.ErrUnauthorized)
			return
		}

		if err := userService.UnfollowUser(r.Context(), followerID.String(), targetID); err != nil {
			utils.RespondError(w, http.StatusInternalServerError, "Failed to unfollow user", utils.ErrInternalServerError)
			return
		}

		utils.RespondJSON(w, http.StatusOK, "User unfollowed", nil)
	}
}

func GetFollowersHandler(userService services.UserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			utils.RespondError(w, http.StatusMethodNotAllowed, "Method not allowed", utils.ErrMethodNotAllowed)
			return
		}

		id := strings.TrimPrefix(r.URL.Path, "/users/followers/")
		//Get the viewer ID from the context
		viewerID, err := utils.GetUserIDFromContext(r.Context())
		if err != nil {
			utils.RespondError(w, http.StatusUnauthorized, "Unauthorized", utils.ErrUnauthorized)
			return
		}
		followers, err := userService.GetFollowers(r.Context(), viewerID.String(), id)
		if err != nil {
			utils.RespondError(w, http.StatusInternalServerError, "Failed to get followers", utils.ErrInternalServerError)
			return
		}
		for _, follower := range followers {
			if follower.ImageURL != nil && *follower.ImageURL != "" {
				follower.ImageURL = utils.GetFullImageURLAvatar(follower.ImageURL)
			}
		}

		utils.RespondJSON(w, http.StatusOK, "Followers retrieved", followers)
	}
}

func GetFollowingHandler(userService services.UserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			utils.RespondError(w, http.StatusMethodNotAllowed, "Method not allowed", utils.ErrMethodNotAllowed)
			return
		}

		id := strings.TrimPrefix(r.URL.Path, "/users/following/")
		//Get the viewer ID from the context
		viewerID, err := utils.GetUserIDFromContext(r.Context())
		if err != nil {
			utils.RespondError(w, http.StatusUnauthorized, "Unauthorized", utils.ErrUnauthorized)
			return
		}
		following, err := userService.GetFollowing(r.Context(), viewerID.String(), id)
		if err != nil {
			utils.RespondError(w, http.StatusInternalServerError, "Failed to get following", utils.ErrInternalServerError)
			return
		}
		for _, follower := range following {
			if follower.ImageURL != nil && *follower.ImageURL != "" {
				follower.ImageURL = utils.GetFullImageURLAvatar(follower.ImageURL)
			}
		}

		utils.RespondJSON(w, http.StatusOK, "Following retrieved", following)
	}
}
