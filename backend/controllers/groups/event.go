package groups

import (
	"mellow/models"
	"mellow/services"
	"mellow/utils"
	"net/http"
	"strings"
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
		if eventTitle == "" {
			utils.RespondError(w, http.StatusBadRequest, "Event title is required", utils.ErrInvalidPayload)
			return
		}
		eventDateStr := r.FormValue("event_date")
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

		err = groupSvc.InsertEvent(r.Context(), &event)
		if err != nil {
			utils.RespondError(w, http.StatusInternalServerError, "Failed to create event "+err.Error(), utils.ErrInternalServerError)
			return
		}

		utils.RespondJSON(w, http.StatusCreated, "Event created successfully", nil)
	}
}

func InsertEventResponse(groupSvc services.GroupService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			utils.RespondError(w, http.StatusMethodNotAllowed, "Method not allowed", utils.ErrMethodNotAllowed)
			return
		}

		userID, err := utils.GetUserIDFromContext(r.Context())
		if err != nil {
			utils.RespondError(w, http.StatusUnauthorized, "Unauthorized", utils.ErrUnauthorized)
			return
		}
		//recupere l'id de l'url
		eventId := strings.TrimPrefix(r.URL.Path, "/groups/events/vote/")
		if eventId == "" {
			utils.RespondError(w, http.StatusBadRequest, "Event ID is required", utils.ErrInvalidPayload)
			return
		}
		// Vérifier si l'ID de l'événement est un UUID valide
		if _, err := uuid.Parse(eventId); err != nil {
			utils.RespondError(w, http.StatusBadRequest, "Invalid event ID format", utils.ErrInvalidPayload)
			return
		}

		vote := r.FormValue("vote")
		if vote == "" {
			utils.RespondError(w, http.StatusBadRequest, "Vote is required", utils.ErrInvalidPayload)
			return
		}

		eventIdValue, err := uuid.Parse(eventId)
		if err != nil {
			utils.RespondError(w, http.StatusBadRequest, "Invalid event ID", utils.ErrInvalidPayload)
			return
		}
		var voteBool bool
		switch vote {
		case "yes":
			voteBool = true
		case "no":
			voteBool = false
		default:
			utils.RespondError(w, http.StatusBadRequest, "Invalid vote value, must be 'yes' or 'no'", utils.ErrInvalidPayload)
			return
		}

		response := models.EventResponse{
			EventID: eventIdValue,
			UserID:  userID,
			Vote:    voteBool,
		}

		err = groupSvc.InsertEventResponse(r.Context(), &response)
		if err != nil {
			utils.RespondError(w, http.StatusInternalServerError, "Failed to insert event response: "+err.Error(), utils.ErrInternalServerError)
			return
		}

		utils.RespondJSON(w, http.StatusOK, "Event response recorded successfully", nil)
	}
}

func GetGroupEvents(groupSvc services.GroupService, groupID string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
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
