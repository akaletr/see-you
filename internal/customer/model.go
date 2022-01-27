package customer

import "github.com/google/uuid"

type customer struct {
	UUID uuid.UUID `json:"uuid"`
	Name string    `json:"name"`
}
