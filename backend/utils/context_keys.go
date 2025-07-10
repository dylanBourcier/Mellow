package utils

import (
	"context"
	"errors"

	"github.com/google/uuid"
)

type ContextKey string

const CtxKeyUserID ContextKey = "user_id"

var ErrNoUserInContext = errors.New("no user ID in context")

func GetUserIDFromContext(ctx context.Context) (uuid.UUID, error) {
	userIDRaw := ctx.Value(CtxKeyUserID)
	if userIDRaw == nil {
		return uuid.Nil, ErrNoUserInContext
	}
	userID, ok := userIDRaw.(uuid.UUID)
	if !ok {
		return uuid.Nil, ErrNoUserInContext
	}
	return userID, nil
}
