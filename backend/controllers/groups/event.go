package groups

import (
	"fmt"
	"mellow/models"
	"mellow/services"
	"mellow/utils"
	"net/http"
	"time"

	"github.com/google/uuid"
)

func InsertGroupEvent(groupSvc services.GroupService, groupID string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseMultipartForm(10 << 20); err != nil {
			utils.RespondError(w, http.StatusBadRequest, "Invalid form data", utils.ErrInvalidPayload)
			return
		}

		userID, err := utils.GetUserIDFromContext(r.Context())
		if err != nil {
			utils.RespondError(w, http.StatusUnauthorized, "Unauthorized", utils.ErrUnauthorized)
			return
		}

		eventTitle := r.FormValue("title")
		fmt.Println("Received event title:", eventTitle)
		if eventTitle == "" {
			utils.RespondError(w, http.StatusBadRequest, "Event title is required", utils.ErrInvalidPayload)
			return
		}
		eventDateStr := r.FormValue("event_date")
		fmt.Println("Received event date string:", eventDateStr)
		var eventDate time.Time
		if eventDateStr != "" {
			// Parse the event date from the form value
			eventDate, err = time.Parse("2006-01-02T15:04", eventDateStr)
			if err != nil {
				utils.RespondError(w, http.StatusBadRequest, "Invalid event date format : "+err.Error(), utils.ErrInvalidPayload)
				return
			}
			if eventDate.Before(time.Now()) {
				utils.RespondError(w, http.StatusBadRequest, "Event date cannot be in the past", utils.ErrInvalidPayload)
				return
			}

		} else {
			utils.RespondError(w, http.StatusBadRequest, "Event date is required", utils.ErrInvalidPayload)
			return
		}

		groupIdValue, err := uuid.Parse(groupID)
		if err != nil {
			utils.RespondError(w, http.StatusBadRequest, "Invalid group ID", utils.ErrInvalidPayload)
			return
		}

		event := models.Event{
			GroupID:   groupIdValue,
			UserID:    userID,
			Title:     eventTitle,
			EventDate: eventDate,
		}
		fmt.Println("event.EventDate:", event.EventDate)
		fmt.Println("Received event date:", event)

		err = groupSvc.InsertEvent(r.Context(), &event)
		if err != nil {
			utils.RespondError(w, http.StatusInternalServerError, "Failed to create event "+err.Error(), utils.ErrInternalServerError)
			return
		}

		utils.RespondJSON(w, http.StatusCreated, "Event created successfully", nil)
	}
}

func GetGroupEvents(groupSvc services.GroupService, groupID string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			utils.RespondError(w, http.StatusMethodNotAllowed, "Method not allowed", utils.ErrMethodNotAllowed)
			return
		}

		if groupID == "" {
			utils.RespondError(w, http.StatusBadRequest, "Group ID is required", utils.ErrInvalidPayload)
			return
		}
		userID, err := utils.GetUserIDFromContext(r.Context())
		if err != nil {
			utils.RespondError(w, http.StatusUnauthorized, "Unauthorized", utils.ErrUnauthorized)
			return
		}
		//Verifier que l'utilisateur a accès au groupe
		isMember, err := groupSvc.IsMember(r.Context(), groupID, userID.String())
		if err != nil {
			utils.RespondError(w, http.StatusInternalServerError, "Failed to check group membership: "+err.Error(), utils.ErrInternalServerError)
			return
		}
		if !isMember {
			utils.RespondJSON(w, http.StatusOK, "Not member", nil)
			return
		}

		events, err := groupSvc.GetGroupEvents(r.Context(), groupID)
		if err != nil {
			utils.RespondError(w, http.StatusInternalServerError, "Failed to get events: "+err.Error(), utils.ErrInternalServerError)
			return
		}

		if len(events) == 0 {
			utils.RespondJSON(w, http.StatusOK, "No events found for this group", nil)
			return
		}

		utils.RespondJSON(w, http.StatusOK, "Events retrieved successfully", events)
	}
}
