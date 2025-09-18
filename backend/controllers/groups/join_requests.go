package groups

import (
	"errors"
	"mellow/services"
	"mellow/utils"
	"net/http"
	"strings"
)

// POST /groups/{id}/join-requests
func CreateJoinRequest(svc services.GroupJoinRequestService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			utils.RespondError(w, http.StatusMethodNotAllowed, "Method not allowed", utils.ErrMethodNotAllowed)
			return
		}
		groupID := strings.TrimPrefix(r.URL.Path, "/groups/")
		groupID = strings.TrimSuffix(groupID, "/join-requests")
		parts := strings.Split(groupID, "/join-requests")
		if strings.Contains(groupID, "/join-requests") {
			groupID = parts[0]
		} else {
			// alternative extraction
			groupID = strings.TrimSuffix(strings.TrimPrefix(r.URL.Path, "/groups/"), "/join-requests")
		}
		if groupID == "" || strings.Contains(groupID, "/") {
			utils.RespondError(w, http.StatusNotFound, "Groupe introuvable", utils.ErrGroupNotFound)
			return
		}
		uid, err := utils.GetUserIDFromContext(r.Context())
		if err != nil {
			utils.RespondError(w, http.StatusUnauthorized, "Unauthorized", utils.ErrUnauthorized)
			return
		}
		req, err := svc.RequestJoin(r.Context(), uid.String(), groupID)
		if err != nil {
			switch {
			case errors.Is(err, utils.ErrGroupNotFound):
				utils.RespondError(w, http.StatusNotFound, "Groupe introuvable", utils.ErrGroupNotFound)
			case errors.Is(err, utils.ErrResourceConflict):
				utils.RespondError(w, http.StatusConflict, "Demande déjà en attente ou déjà membre", utils.ErrResourceConflict)
			case errors.Is(err, utils.ErrInvalidPayload):
				utils.RespondError(w, http.StatusBadRequest, "Données invalides", utils.ErrInvalidPayload)
			default:
				utils.RespondError(w, http.StatusInternalServerError, "Erreur interne", utils.ErrInternalServerError)
			}
			return
		}
		res := map[string]interface{}{
			"id":          req.ID,
			"groupId":     req.GroupID,
			"requesterId": req.RequesterID,
			"status":      req.Status,
			"createdAt":   req.CreatedAt,
		}
		utils.RespondJSON(w, http.StatusCreated, "Join request created", res)
	}
}

// GET /groups/{id}/join-requests
func ListJoinRequests(svc services.GroupJoinRequestService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			utils.RespondError(w, http.StatusMethodNotAllowed, "Method not allowed", utils.ErrMethodNotAllowed)
			return
		}
		groupID := strings.TrimPrefix(r.URL.Path, "/groups/")
		groupID = strings.TrimSuffix(groupID, "/join-requests")
		if groupID == "" || strings.Contains(groupID, "/") {
			utils.RespondError(w, http.StatusNotFound, "Groupe introuvable", utils.ErrGroupNotFound)
			return
		}
		uid, err := utils.GetUserIDFromContext(r.Context())
		if err != nil {
			utils.RespondError(w, http.StatusUnauthorized, "Unauthorized", utils.ErrUnauthorized)
			return
		}
		list, err := svc.ListPending(r.Context(), uid.String(), groupID)
		if err != nil {
			switch {
			case errors.Is(err, utils.ErrGroupNotFound):
				utils.RespondError(w, http.StatusNotFound, "Groupe introuvable", utils.ErrGroupNotFound)
			case errors.Is(err, utils.ErrForbidden):
				utils.RespondError(w, http.StatusForbidden, "Forbidden", utils.ErrForbidden)
			default:
				utils.RespondError(w, http.StatusInternalServerError, "Erreur interne", utils.ErrInternalServerError)
			}
			return
		}
		var resp []map[string]interface{}
		for _, it := range list {
			resp = append(resp, map[string]interface{}{
				"id": it.ID,
				"requester": map[string]interface{}{
					"id":       it.RequesterID,
					"username": it.RequesterUsername,
					"avatar":   it.RequesterAvatar,
				},
				"status":    it.Status,
				"createdAt": it.CreatedAt,
			})
		}
		utils.RespondJSON(w, http.StatusOK, "Pending join requests", resp)
	}
}

// POST /groups/{id}/join-requests/{rid}/accept
func AcceptJoinRequest(svc services.GroupJoinRequestService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			utils.RespondError(w, http.StatusMethodNotAllowed, "Method not allowed", utils.ErrMethodNotAllowed)
			return
		}
		path := strings.TrimPrefix(r.URL.Path, "/groups/")
		// path format: {id}/join-requests/{rid}/accept
		parts := strings.Split(path, "/")
		if len(parts) < 4 || parts[1] != "join-requests" || parts[3] != "accept" {
			utils.RespondError(w, http.StatusNotFound, "Not found", utils.ErrBadRequest)
			return
		}
		groupID := parts[0]
		requestID := parts[2]
		uid, err := utils.GetUserIDFromContext(r.Context())
		if err != nil {
			utils.RespondError(w, http.StatusUnauthorized, "Unauthorized", utils.ErrUnauthorized)
			return
		}
		req, err := svc.Accept(r.Context(), uid.String(), groupID, requestID)
		if err != nil {
			switch {
			case errors.Is(err, utils.ErrGroupNotFound):
				utils.RespondError(w, http.StatusNotFound, "Groupe introuvable", utils.ErrGroupNotFound)
			case errors.Is(err, utils.ErrForbidden):
				utils.RespondError(w, http.StatusForbidden, "Forbidden", utils.ErrForbidden)
			case errors.Is(err, utils.ErrResourceConflict):
				utils.RespondError(w, http.StatusConflict, "Conflit d'état", utils.ErrResourceConflict)
			default:
				utils.RespondError(w, http.StatusInternalServerError, "Erreur interne", utils.ErrInternalServerError)
			}
			return
		}
		utils.RespondJSON(w, http.StatusOK, "Join request accepted", map[string]interface{}{
			"status":    req.Status,
			"decidedAt": req.DecidedAt,
			"decidedBy": req.DecidedBy,
		})
	}
}

// POST /groups/{id}/join-requests/{rid}/reject
func RejectJoinRequest(svc services.GroupJoinRequestService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			utils.RespondError(w, http.StatusMethodNotAllowed, "Method not allowed", utils.ErrMethodNotAllowed)
			return
		}
		path := strings.TrimPrefix(r.URL.Path, "/groups/")
		parts := strings.Split(path, "/")
		if len(parts) < 4 || parts[1] != "join-requests" || parts[3] != "reject" {
			utils.RespondError(w, http.StatusNotFound, "Not found", utils.ErrBadRequest)
			return
		}
		groupID := parts[0]
		requestID := parts[2]
		uid, err := utils.GetUserIDFromContext(r.Context())
		if err != nil {
			utils.RespondError(w, http.StatusUnauthorized, "Unauthorized", utils.ErrUnauthorized)
			return
		}
		req, err := svc.Reject(r.Context(), uid.String(), groupID, requestID)
		if err != nil {
			switch {
			case errors.Is(err, utils.ErrGroupNotFound):
				utils.RespondError(w, http.StatusNotFound, "Groupe introuvable", utils.ErrGroupNotFound)
			case errors.Is(err, utils.ErrForbidden):
				utils.RespondError(w, http.StatusForbidden, "Forbidden", utils.ErrForbidden)
			case errors.Is(err, utils.ErrResourceConflict):
				utils.RespondError(w, http.StatusConflict, "Conflit d'état", utils.ErrResourceConflict)
			default:
				utils.RespondError(w, http.StatusInternalServerError, "Erreur interne", utils.ErrInternalServerError)
			}
			return
		}
		utils.RespondJSON(w, http.StatusOK, "Join request rejected", map[string]interface{}{
			"status":    req.Status,
			"decidedAt": req.DecidedAt,
			"decidedBy": req.DecidedBy,
		})
	}
}

// DELETE /groups/{id}/join-requests/self or GET self status
func SelfJoinRequestRouter(svc services.GroupJoinRequestService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		groupID := strings.TrimPrefix(r.URL.Path, "/groups/")
		groupID = strings.TrimSuffix(groupID, "/join-requests/self")
		if groupID == "" || strings.Contains(groupID, "/") {
			utils.RespondError(w, http.StatusNotFound, "Groupe introuvable", utils.ErrGroupNotFound)
			return
		}
		uid, err := utils.GetUserIDFromContext(r.Context())
		if err != nil {
			utils.RespondError(w, http.StatusUnauthorized, "Unauthorized", utils.ErrUnauthorized)
			return
		}
		switch r.Method {
		case http.MethodDelete:
			if err := svc.Cancel(r.Context(), uid.String(), groupID); err != nil {
				switch {
				case errors.Is(err, utils.ErrBadRequest):
					utils.RespondError(w, http.StatusBadRequest, "No pending request", utils.ErrBadRequest)
				default:
					utils.RespondError(w, http.StatusInternalServerError, "Erreur interne", utils.ErrInternalServerError)
				}
				return
			}
			w.WriteHeader(http.StatusNoContent)
		case http.MethodGet:
			req, err := svc.GetSelfPending(r.Context(), uid.String(), groupID)
			if err != nil {
				utils.RespondError(w, http.StatusInternalServerError, "Erreur interne", utils.ErrInternalServerError)
				return
			}
			if req == nil {
				utils.RespondJSON(w, http.StatusOK, "No pending request", nil)
				return
			}
			utils.RespondJSON(w, http.StatusOK, "Pending join request", map[string]interface{}{
				"id":        req.ID,
				"status":    req.Status,
				"createdAt": req.CreatedAt,
			})
		default:
			utils.RespondError(w, http.StatusMethodNotAllowed, "Method not allowed", utils.ErrMethodNotAllowed)
		}
	}
}
