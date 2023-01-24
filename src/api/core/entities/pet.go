package entities

import "github.com/rafawilliner/sama-api/src/api/core/entities/constants"

type Pet struct {
	Id     int64
	Name   string
	Gender string
	Race   string
	Age    int32
	Specie constants.Specie
	Weight int32
}

func NewPet(name string, gender string, race string, age int32, specie constants.Specie, weight int32) *Pet {
	return &Pet{
		Name:   name,
		Gender: gender,
		Race:   race,
		Age:    age,
		Specie: specie,
		Weight: weight,
	}
}
