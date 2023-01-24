package create

import (
	"context"

	"github.com/rafawilliner/sama-api/src/api/core/contracts/create"
	"github.com/rafawilliner/sama-api/src/api/core/entities"
	"github.com/rafawilliner/sama-api/src/api/core/providers"
)

type PetUseCase interface {
	Execute(ctx context.Context, request *create.PetRequest) (*entities.Pet, error)
}

type PetUseCaseImpl struct {
	PetProvider providers.Pet
}

func (useCase *PetUseCaseImpl) Execute(ctx context.Context, request *create.PetRequest) (*entities.Pet, error) {

	pet := entities.NewPet(request.Name, request.Gender, request.Race, request.Age, request.Specie, request.Weight)
	err := useCase.PetProvider.Save(ctx, pet)
	if err != nil {
		return nil, err
	}

	return pet, nil
}
