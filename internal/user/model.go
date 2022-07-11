package user

import (
	"math/rand"
)

type User struct {
	ID   uint64 `json:"id,omitempty"`
	Name string `json:"name"`
	Age  uint8  `json:"age"`
}

func (u *User) Save() (uint64, error) {
	u.ID = rand.Uint64()
	return u.ID, nil
}
