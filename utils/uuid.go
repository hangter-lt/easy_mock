package utils

import (
	"strings"

	"github.com/google/uuid"
)

func UUID() string {
	uid := uuid.New().String()
	return strings.ReplaceAll(uid, "-", "")
}
