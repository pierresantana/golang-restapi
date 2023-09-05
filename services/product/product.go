package product

import (
	"context"

	"github.com/pierresantana/golang-restapi/models"
	"github.com/pierresantana/golang-restapi/repositories/product"
)

type ProductsService struct {
	r product.ProductRepository
}

func NewService(r product.ProductRepository) *ProductsService {
	return &ProductsService{
		r: r,
	}
}

func (m *ProductsService) GetAll(ctx context.Context) ([]*models.Product, error) {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	return m.r.GetAll(ctx)
}

func (m *ProductsService) GetByID(ctx context.Context, id string) (*models.Product, error) {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	return m.r.GetByID(ctx, id)
}

func (m *ProductsService) Create(ctx context.Context, product *models.Product) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	_, err := m.r.Insert(ctx, product)
	return err
}

func (m *ProductsService) Delete(ctx context.Context, id string) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	return m.r.Delete(ctx, id)
}

func (m *ProductsService) Update(ctx context.Context, id string, product *models.Product) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	_, err := m.r.Update(ctx, id, product)
	return err
}
