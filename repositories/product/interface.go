package product

import (
	"context"

	"github.com/pierresantana/golang-restapi/models"
)

type ProductRepository interface {
	GetAll(ctx context.Context) ([]*models.Product, error)
	GetByID(ctx context.Context, id string) (*models.Product, error)
	Insert(ctx context.Context, product *models.Product) (*models.Product, error)
	Update(ctx context.Context, id string, product *models.Product) (*models.Product, error)
	Delete(ctx context.Context, id string) error
}
