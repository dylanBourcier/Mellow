// errors.go
package utils

import "errors"

var (
	// Requêtes mal formées ou invalides
	ErrInvalidPayload  = errors.New("INVALID_PAYLOAD")
	ErrBadRequest      = errors.New("BAD_REQUEST")
	ErrInvalidJSON     = errors.New("INVALID_JSON_FORMAT")
	ErrInvalidUserData = errors.New("INVALID_USER_DATA")

	// Authentification / Autorisation
	ErrUnauthorized       = errors.New("UNAUTHORIZED")
	ErrForbidden          = errors.New("FORBIDDEN")
	ErrInvalidCredentials = errors.New("INVALID_CREDENTIALS")
	ErrSessionExpired     = errors.New("SESSION_EXPIRED")

	// Méthodes HTTP non autorisées
	ErrMethodNotAllowed = errors.New("METHOD_NOT_ALLOWED")

	// Ressources manquantes
	ErrUserNotFound    = errors.New("USER_NOT_FOUND")
	ErrPostNotFound    = errors.New("POST_NOT_FOUND")
	ErrGroupNotFound   = errors.New("GROUP_NOT_FOUND")
	ErrMessageNotFound = errors.New("MESSAGE_NOT_FOUND")
	ErrSessionNotFound = errors.New("SESSION_NOT_FOUND")

	// Conflits
	ErrUsernameAlreadyExists = errors.New("USERNAME_ALREADY_EXISTS")
	ErrEmailAlreadyExists    = errors.New("EMAIL_ALREADY_EXISTS")
	ErrUserAlreadyExists     = errors.New("USER_ALREADY_EXISTS")
	ErrResourceConflict      = errors.New("RESOURCE_CONFLICT")

	// Limites et contraintes
	ErrTooManyRequests = errors.New("TOO_MANY_REQUESTS")

	// Autres
	ErrInternalServerError = errors.New("INTERNAL_SERVER_ERROR")
	ErrNotImplemented      = errors.New("NOT_IMPLEMENTED")
	ErrUUIDGeneration	  = errors.New("UUID_GENERATION_FAILED")
)
