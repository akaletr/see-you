package storage

import "cmd/main/main.go/internal/user"

type Storage interface {
	Start() error

	Read(id string) (user.User, error)
	Write(id string, user user.User) error
}
