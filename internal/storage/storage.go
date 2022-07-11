package storage

import (
	"cmd/main/main.go/internal/user"
	"errors"
)

type storage struct {
	db map[string]user.User
}

func New() Storage {
	return &storage{db: map[string]user.User{}}
}

func (s *storage) Read(id string) (user.User, error) {
	if u, ok := s.db[id]; ok {
		return u, nil
	}
	return user.User{}, errors.New("no user with id " + id)
}

func (s *storage) Write(id string, user user.User) error {
	s.db[id] = user
	return nil
}
