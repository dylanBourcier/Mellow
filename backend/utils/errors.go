// errors.go
package utils

const (
	// Requêtes mal formées ou invalides
	ErrInvalidPayload  = "INVALID_PAYLOAD"
	ErrBadRequest      = "BAD_REQUEST"
	ErrInvalidJSON     = "INVALID_JSON_FORMAT"
	ErrInvalidUserData = "INVALID_USER_DATA"

	// Authentification / Autorisation
	ErrUnauthorized       = "UNAUTHORIZED"
	ErrForbidden          = "FORBIDDEN"
	ErrInvalidCredentials = "INVALID_CREDENTIALS"
	ErrSessionExpired     = "SESSION_EXPIRED"

	// Méthodes HTTP non autorisées
	ErrMethodNotAllowed = "METHOD_NOT_ALLOWED"

	// Ressources manquantes
	ErrUserNotFound    = "USER_NOT_FOUND"
	ErrPostNotFound    = "POST_NOT_FOUND"
	ErrGroupNotFound   = "GROUP_NOT_FOUND"
	ErrMessageNotFound = "MESSAGE_NOT_FOUND"

	// Conflits
	ErrUsernameAlreadyExists = "USERNAME_ALREADY_EXISTS"
	ErrEmailAlreadyExists    = "EMAIL_ALREADY_EXISTS"
	ErrUserAlreadyExists     = "USER_ALREADY_EXISTS"
	ErrResourceConflict      = "RESOURCE_CONFLICT"

	// Limites et contraintes
	ErrTooManyRequests = "TOO_MANY_REQUESTS"

	// Autres
	ErrInternalServerError = "INTERNAL_SERVER_ERROR"
	ErrNotImplemented      = "NOT_IMPLEMENTED"
)
