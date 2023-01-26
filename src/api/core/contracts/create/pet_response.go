package create

import (
	"github.com/rafawilliner/sama-api/src/api/core/entities"
	"github.com/rafawilliner/sama-api/src/api/core/entities/constants"
)

type PetResponse struct {
	Id     int64            `json:"id"`
	Name   string           `json:"description"`
	Gender *string          `json:"family"`
	Race   *string          `json:"race"`
	Age    *int32           `json:"age"`
	Specie constants.Specie `json:"specie"`
	Weight *int32           `json:"weight"`
}

func NewResponse(pet *entities.Pet) PetResponse {
	response := PetResponse{
		Id:     pet.Id,
		Name:   pet.Name,
		Gender: pet.Gender,
		Race:   pet.Race,
		Age:    pet.Age,
		Specie: pet.Specie,
		Weight: pet.Weight,
	}

	return response
}
