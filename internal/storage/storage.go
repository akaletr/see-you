package storage

import (
	"cmd/main/main.go/internal/user"
)

type storage struct {
	db map[string]user.User
}

func New() Storage {
	return nil
}

func (s *storage) Read(id string) (user.User, error) {
	return user.User{}, nil
}

func (s *storage) Write(id string, user user.User) error {
	return nil
}
