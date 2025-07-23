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
		if err := r.ParseForm(); err != nil {
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
		var eventDate time.Time
		if eventDateStr != "" {
			// Parse the event date from the form value
			eventDate, err = time.Parse(time.RFC3339, eventDateStr)
			if err != nil {
				utils.RespondError(w, http.StatusBadRequest, "Invalid event date format", utils.ErrInvalidPayload)
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
