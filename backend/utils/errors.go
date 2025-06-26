// errors.go
package utils

const (
	// Requêtes mal formées ou invalides
	ErrInvalidPayload = "INVALID_PAYLOAD"
	ErrBadRequest     = "BAD_REQUEST"

	// Authentification / Autorisation
	ErrUnauthorized       = "UNAUTHORIZED"
	ErrForbidden          = "FORBIDDEN"
	ErrInvalidCredentials = "INVALID_CREDENTIALS"
	ErrSessionExpired     = "SESSION_EXPIRED"

	// Ressources manquantes
	ErrUserNotFound    = "USER_NOT_FOUND"
	ErrPostNotFound    = "POST_NOT_FOUND"
	ErrGroupNotFound   = "GROUP_NOT_FOUND"
	ErrMessageNotFound = "MESSAGE_NOT_FOUND"

	// Conflits
	ErrUsernameAlreadyExists = "USERNAME_ALREADY_EXISTS"
	ErrEmailAlreadyExists    = "EMAIL_ALREADY_EXISTS"
	ErrResourceConflict      = "RESOURCE_CONFLICT"

	// Limites et contraintes
	ErrTooManyRequests = "TOO_MANY_REQUESTS"

	// Autres
	ErrInternalServerError = "INTERNAL_SERVER_ERROR"
	ErrNotImplemented      = "NOT_IMPLEMENTED"
)
