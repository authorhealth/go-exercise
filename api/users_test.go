package api

import (
	"context"
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/authorhealth/go-exercise/domain"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestGetUsers(t *testing.T) {
	assert := assert.New(t)

	rr := httptest.NewRecorder()
	r := httptest.NewRequest("", "/", nil)

	entities := []*domain.User{
		{
			ID:        uuid.New().String(),
			NameFirst: "John",
			NameLast:  "Smith",
			Email:     "jsmith@gmail.com",
		},
		{
			ID:        uuid.New().String(),
			NameFirst: "Jane",
			NameLast:  "Doe",
			Email:     "jane.doe@gmail.com",
		},
	}

	expectedRes := &getUsersResponse{
		Users: []*user{
			{
				ID:        entities[0].ID,
				NameFirst: entities[0].NameFirst,
				NameLast:  entities[0].NameLast,
				Email:     entities[0].Email,
			},
			{
				ID:        entities[1].ID,
				NameFirst: entities[1].NameFirst,
				NameLast:  entities[1].NameLast,
				Email:     entities[1].Email,
			},
		},
	}

	store := domain.NewMockStorer(t)
	userStore := domain.NewMockUserStorer(t)

	store.On("Users").Return(userStore).Once()
	userStore.On("Find", r.Context()).Return(entities, nil).Once()

	GetUsers(store)(rr, r)

	res := &getUsersResponse{}
	err := json.Unmarshal(rr.Body.Bytes(), res)
	assert.NoError(err)

	assert.Equal(expectedRes, res)
}

func TestGetUserByID(t *testing.T) {
	assert := assert.New(t)

	rr := httptest.NewRecorder()
	r := httptest.NewRequest("", "/", nil)

	entity := &domain.User{
		ID:        uuid.New().String(),
		NameFirst: "John",
		NameLast:  "Smith",
		Email:     "jsmith@gmail.com",
	}

	expectedRes := &user{
		ID:        entity.ID,
		NameFirst: entity.NameFirst,
		NameLast:  entity.NameLast,
		Email:     entity.Email,
	}

	r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, &chi.Context{
		URLParams: chi.RouteParams{
			Keys:   []string{"id"},
			Values: []string{entity.ID},
		},
	}))

	store := domain.NewMockStorer(t)
	userStore := domain.NewMockUserStorer(t)

	store.On("Users").Return(userStore).Once()
	userStore.On("FindByID", r.Context(), entity.ID).Return(entity, nil).Once()

	GetUserByID(store)(rr, r)

	res := &user{}
	err := json.Unmarshal(rr.Body.Bytes(), res)
	assert.NoError(err)

	assert.Equal(expectedRes, res)
}
