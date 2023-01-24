package providers

import (
	"context"

	"github.com/rafawilliner/sama-api/src/api/core/entities"
)

type Pet interface {
	Save(ctx context.Context, pet *entities.Pet) error
}
