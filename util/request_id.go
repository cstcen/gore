package util

import (
	"github.com/google/uuid"
	"strings"
)

func GetRequestID() string {
	return strings.ReplaceAll(uuid.New().String(), "-", "")
}
