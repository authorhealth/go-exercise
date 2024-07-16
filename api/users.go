package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/authorhealth/go-exercise/domain"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

type createUserRequest struct {
	NameFirst string `json:"nameFirst"`
	NameLast  string `json:"nameLast"`
	Email     string `json:"email"`
}

func CreateUser(store domain.Storer) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		req := &createUserRequest{}
		err := json.NewDecoder(r.Body).Decode(req)
		if err != nil {
			respondError(w, r, http.StatusBadRequest, fmt.Errorf("decoding request body: %w", err))
			return
		}

		user := &domain.User{
			ID:        uuid.New().String(),
			NameFirst: req.NameFirst,
			NameLast:  req.NameLast,
			Email:     req.Email,
		}

		err = store.Users().Save(ctx, user)
		if err != nil {
			respondError(w, r, http.StatusBadRequest, fmt.Errorf("saving user: %w", err))
			return
		}

		respondJSON(w, http.StatusCreated, nil)
	}
}

type user struct {
	ID        string `json:"id"`
	NameFirst string `json:"nameFirst"`
	NameLast  string `json:"nameLast"`
	Email     string `json:"email"`
}

type getUsersResponse struct {
	Users []*user `json:"users"`
}

func GetUsers(store domain.Storer) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		users, err := store.Users().Find(ctx)
		if err != nil {
			respondError(w, r, http.StatusInternalServerError, fmt.Errorf("finding users: %w", err))
			return
		}

		res := &getUsersResponse{}

		for _, entity := range users {
			res.Users = append(res.Users, &user{
				ID:        entity.ID,
				NameFirst: entity.NameFirst,
				NameLast:  entity.NameLast,
				Email:     entity.Email,
			})
		}

		respondJSON(w, http.StatusOK, res)
	}
}

func GetUserByID(store domain.Storer) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		id := chi.URLParam(r, "id")

		entity, err := store.Users().FindByID(ctx, id)
		if err != nil {
			if errors.Is(err, domain.ErrNotFound) {
				respondError(w, r, http.StatusNotFound, err)
				return
			}

			respondError(w, r, http.StatusInternalServerError, fmt.Errorf("finding users: %w", err))
			return
		}

		user := &user{
			ID:        entity.ID,
			NameFirst: entity.NameFirst,
			NameLast:  entity.NameLast,
			Email:     entity.Email,
		}

		respondJSON(w, http.StatusOK, user)
	}
}
