package pet

import (
	"context"

	"github.com/rafawilliner/sama-api/src/api/core/entities"
	"gorm.io/gorm"
)

type Repository struct {
	StoreClient *gorm.DB
}

func (repository Repository) Save(ctx context.Context, commercialProduct *entities.Pet) error {

	var err error
	err = repository.StoreClient.Create(commercialProduct).Error
	if err != nil {
		return err
	}

	return nil
}
