package helpers

import (
	"github.com/google/uuid"
)

func GenerateShortCode() string {
	code := uuid.New()
	return code.String()[:6]
}
