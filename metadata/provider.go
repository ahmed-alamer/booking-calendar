package metadata

import "github.com/google/uuid"

type Provider struct {
	Id       uuid.UUID `json:"id"`
	Name     string    `json:"name"`
	Category string    `json:"category"`
}
