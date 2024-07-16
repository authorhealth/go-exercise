package domain

import (
	"context"
	"errors"
)

var ErrNotFound = errors.New("not found")

type Storer interface {
	Users() UserStorer
}

type UserStorer interface {
	Find(context.Context) ([]*User, error)
	FindByID(context.Context, string) (*User, error)
	Save(context.Context, *User) error
}
