package idutil

import (
	"github.com/google/uuid"
)

func GenerateUUID(prefix string) string {
	uuid, _ := uuid.NewRandom()
	return prefix + "-" + uuid.String()
}
