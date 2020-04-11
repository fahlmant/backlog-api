package types

import (
	"github.com/google/uuid"
)

type Game struct {
	ID uuid.UUID `json:"ID"`
	Title string `json:"Title"`
	Platform string `json:"Platform"s`
}
