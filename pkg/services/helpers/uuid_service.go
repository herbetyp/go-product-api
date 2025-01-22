package utils

import "github.com/google/uuid"

func NewUUID() string {
	return uuid.Must(uuid.NewRandom()).String()
}

func UUIDValidate(uuidString string) bool {
	_, err := uuid.Parse(uuidString)
	return err == nil
}
