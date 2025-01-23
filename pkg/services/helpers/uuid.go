package helpers

import "github.com/google/uuid"

func NewUUID() string {
	return uuid.Must(uuid.NewRandom()).String()
}
