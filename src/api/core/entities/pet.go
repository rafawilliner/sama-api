package entities

import "github.com/rafawilliner/sama-api/src/api/core/entities/constants"

type Pet struct {
	Id     int64            `gorm:"PRIMARY_KEY;AUTO_INCREMENT"`
	Name   string           `gorm:"type:varchar(255);not null"`
	Gender *string          `gorm:"type:varchar(100);null"`
	Race   *string          `gorm:"type:varchar(100);null"`
	Age    *int32           `gorm:"type:integer;null"`
	Specie constants.Specie `gorm:"type:varchar(100);not null"`
	Weight *int32           `gorm:"type:integer;null"`
}

func NewPet(name string, gender *string, race *string, age *int32, specie constants.Specie, weight *int32) *Pet {
	return &Pet{
		Name:   name,
		Gender: gender,
		Race:   race,
		Age:    age,
		Specie: specie,
		Weight: weight,
	}
}
