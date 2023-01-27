package mocks

import (
	"context"

	"github.com/rafawilliner/sama-api/src/api/core/contracts/create"
	"github.com/rafawilliner/sama-api/src/api/core/entities"
	"github.com/stretchr/testify/mock"
)

type MockCreatePetUseCase struct {
	mock.Mock
}

func (mock *MockCreatePetUseCase) Execute(ctx context.Context, request *create.PetRequest) (*entities.Pet, error) {
	responseArgs := mock.Called(ctx, request)
	resp := responseArgs.Get(0)
	err := responseArgs.Error(1)
	if resp != nil {
		return resp.(*entities.Pet), err
	}
	
	return nil, err
}