package place

import "github.com/google/uuid"

type place struct {
	UUID uuid.UUID `json:"uuid"`
	Name string    `json:"name"`
}

func NewPlace() {
	return
}
