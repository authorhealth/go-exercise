package memory

import (
	"testing"

	"github.com/authorhealth/go-exercise/domain"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestToModel(t *testing.T) {
	assert := assert.New(t)

	entity := &domain.User{
		ID:        uuid.New().String(),
		NameFirst: "John",
		NameLast:  "Smith",
		Email:     "john.smith@gmail.com",
	}

	model, err := toModel(entity)
	assert.NoError(err)

	assert.Equal(&user{
		ID:        entity.ID,
		NameFirst: entity.NameFirst,
		NameLast:  entity.NameLast,
		Email:     entity.Email,
	}, model)
}

func TestToEntity(t *testing.T) {
	assert := assert.New(t)

	model := &user{
		ID:        uuid.New().String(),
		NameFirst: "John",
		NameLast:  "Smith",
		Email:     "john.smith@gmail.com",
	}

	entity, err := toEntity(model)
	assert.NoError(err)

	assert.Equal(&domain.User{
		ID:        entity.ID,
		NameFirst: entity.NameFirst,
		NameLast:  entity.NameLast,
		Email:     entity.Email,
	}, entity)
}
