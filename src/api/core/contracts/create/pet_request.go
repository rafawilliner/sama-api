package create

import (
	"github.com/rafawilliner/sama-api/src/api/core/entities/constants"
	"github.com/rafawilliner/sama-api/src/api/core/errors"
)

type PetRequest struct {
	Id     int64            `json:"id"`
	Name   string           `json:"name" binding:"required,max=255"`
	Gender *string          `json:"family"`
	Race   *string          `json:"race"`
	Age    *int32           `json:"age"`
	Specie constants.Specie `json:"specie"`
	Weight *int32           `json:"weight"`
}

func (request *PetRequest) ValidateRequest() error {
	if request.Age != nil && *request.Age <= 0 {
		err := errors.NewValidationError("Age must be greater than 0")
		return err
	}
	return nil
}

func NewRequest() PetRequest {
	return PetRequest{}
}
