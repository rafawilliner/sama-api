package dependencies

import (
	"github.com/rafawilliner/sama-api/src/api/config/database"
	"github.com/rafawilliner/sama-api/src/api/core/usecases/create"
	"github.com/rafawilliner/sama-api/src/api/entrypoints"
	createHandler "github.com/rafawilliner/sama-api/src/api/entrypoints/handlers/create"
	petRepo "github.com/rafawilliner/sama-api/src/api/repositories/pet"
)

type HandlerContainer struct {
	PetCreate entrypoints.Handler
}

type StartConnection struct {
	StoreConnection database.Connection
}

func (connections StartConnection) Start() *HandlerContainer {

	// Database
	storeClient, err := connections.StoreConnection.Connect()
	if err != nil {
		panic(err.Error())
	}

	// Providers
	petProvider := petRepo.Repository{
		StoreClient: storeClient,
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
