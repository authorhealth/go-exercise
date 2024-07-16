package memory

import (
	"context"
	"errors"
	"fmt"
	"sync"

	"github.com/authorhealth/go-exercise/domain"
)

type UserStore struct {
	lock  sync.Mutex
	users map[string]*user
}

var _ domain.UserStorer = (*UserStore)(nil)

func NewUserStore() *UserStore {
	return &UserStore{
		users: make(map[string]*user),
	}
}

type user struct {
	ID        string
	NameFirst string
	NameLast  string
	Email     string
}

func toModel(entity *domain.User) (*user, error) {
	if len(entity.ID) == 0 {
		return nil, errors.New("entity ID field is empty")
	}

	return &user{
		ID:        entity.ID,
		NameFirst: entity.NameFirst,
		NameLast:  entity.NameLast,
		Email:     entity.Email,
	}, nil
}

func toEntity(model *user) (*domain.User, error) {
	return &domain.User{
		ID:        model.ID,
		NameFirst: model.NameFirst,
		NameLast:  model.NameLast,
		Email:     model.Email,
	}, nil
}

func (u *UserStore) Find(ctx context.Context) ([]*domain.User, error) {
	u.lock.Lock()
	defer u.lock.Unlock()

	var entities []*domain.User

	for _, model := range u.users {
		entity, err := toEntity(model)
		if err != nil {
			return nil, fmt.Errorf("to entity: %w", err)
		}

		entities = append(entities, entity)
	}

	return entities, nil
}

func (u *UserStore) FindByID(ctx context.Context, id string) (*domain.User, error) {
	u.lock.Lock()
	defer u.lock.Unlock()

	model, ok := u.users[id]
	if !ok {
		return nil, domain.ErrNotFound
	}

	entity, err := toEntity(model)
	if err != nil {
		return nil, fmt.Errorf("to entity: %w", err)
	}

	return entity, nil
}

func (u *UserStore) Save(ctx context.Context, user *domain.User) error {
	u.lock.Lock()
	defer u.lock.Unlock()

	model, err := toModel(user)
	if err != nil {
		return fmt.Errorf("to model: %w", err)
	}

	u.users[model.ID] = model

	return nil
}
