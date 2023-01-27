package mocks

import (
	"context"

	"github.com/rafawilliner/sama-api/src/api/core/entities"
	"github.com/stretchr/testify/mock"
)

type PetProvider struct {
	mock.Mock
}

func (mock *PetProvider) Save(ctx context.Context, pet *entities.Pet) error {
	responseArgs := mock.Called(ctx, pet)
	err := responseArgs.Error(0)

	return err
}
