package util

import (
	"context"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"strings"
)

const RequestIDContextKey = "X-Request-ID"

func MustRequestID(c context.Context) string {
	id, ok := c.Value(RequestIDContextKey).(string)
	if !ok {
		panic(errors.New(RequestIDContextKey + " not found"))
	}
	return id
}

func GenerateRequestID() string {
	return strings.ReplaceAll(uuid.New().String(), "-", "")
}
