package dependencies

import (
	"github.com/rafawilliner/sama-api/src/api/core/usecases/create"
	"github.com/rafawilliner/sama-api/src/api/entrypoints"
	createHandler "github.com/rafawilliner/sama-api/src/api/entrypoints/handlers/create"
	petRepo "github.com/rafawilliner/sama-api/src/api/repositories/pet"
)

type HandlerContainer struct {
	PetCreate entrypoints.Handler
}

func Start() *HandlerContainer {

	// Providers
	petProvider := petRepo.Repository{
		//TODO connection
	}

	// UseCases
	petUseCase := &create.PetUseCaseImpl{
		PetProvider: petProvider,
	}

	//Handlers
	handlers := HandlerContainer{}
	handlers.PetCreate = &createHandler.Pet{
		CreateUseCase: petUseCase,
	}

	return &handlers
}
