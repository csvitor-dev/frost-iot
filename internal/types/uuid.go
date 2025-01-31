package types

import "github.com/google/uuid"

// represents a UUID version
type UUID string

func NewUUID() UUID {
	return UUID(uuid.NewString())
}