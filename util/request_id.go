package util

import (
	"context"
	"github.com/google/uuid"
	"strings"
)

const RequestIDContextKey = "X-Request-ID"

func MustRequestID(c context.Context) string {
	id, ok := c.Value(RequestIDContextKey).(string)
	if !ok {
		return ""
	}
	return id
}

func GenerateRequestID() string {
	return strings.ReplaceAll(uuid.New().String(), "-", "")
}
