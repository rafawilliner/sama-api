package create

import (
	"context"
	"testing"

	"github.com/rafawilliner/sama-api/src/api/core/contracts/create"
	"github.com/rafawilliner/sama-api/src/api/core/entities"
	"github.com/rafawilliner/sama-api/src/api/core/errors"
	"github.com/rafawilliner/sama-api/src/api/core/providers/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	mockPetProvider *mocks.PetProvider
	mockedRequest   = getMockedRequest()
	mockCtx         = func(ctx context.Context) bool {
		return true
	}
)

func TestCreatingPetWhenIsSuccessShouldReturnPet(t *testing.T) {
	// Given
	ctx := context.Background()
	initMockedDependencies()
	useCase := getUseCase()
	pet := getEntity()

	mockPetProvider.On("Save", mock.MatchedBy(mockCtx), pet).
	Return(nil).
		Run(func(args mock.Arguments) {
			pet := args.Get(1).(*entities.Pet)
			assert.Equal(t, mockedRequest.Name, pet.Name)
		})

	// When
	response, err := useCase.Execute(ctx, mockedRequest)

	// Then
	assert.NotNil(t, response)
	assert.NoError(t, err)
	mockPetProvider.AssertNumberOfCalls(t, "Save", 1)
}

func TestCreatingPetWhenBadRequestShouldReturnError(t *testing.T) {
	// Given
	ctx := context.Background()
	initMockedDependencies()
	useCase := getUseCase()
	pet := getEntity()
	repositoryError := errors.NewRepositoryError("Error saving resource")

	mockPetProvider.On("Save", mock.MatchedBy(mockCtx), pet).Return(repositoryError).
		Run(func(args mock.Arguments) {
			pet := args.Get(1).(*entities.Pet)
			assert.Equal(t, mockedRequest.Name, pet.Name)
		})

	// When
	response, err := useCase.Execute(ctx, mockedRequest)

	// Then
	assert.Nil(t, response)
	assert.Error(t, err)
	mockPetProvider.AssertNumberOfCalls(t, "Save", 1)
	assert.Equal(t, "Error saving resource", err.Error())
}

func getUseCase() PetUseCaseImpl {
	return PetUseCaseImpl{
		PetProvider: mockPetProvider,
	}
}

func initMockedDependencies() {
	mockPetProvider = new(mocks.PetProvider)
}

func getMockedRequest() *create.PetRequest {
	return &create.PetRequest{
		Name: "Firu",
	}
}

func getEntity() *entities.Pet {
	return &entities.Pet{
		Name: "Firu",
	}
}
