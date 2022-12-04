package usecase

import (
	"v-commerce-service/models"
	"v-commerce-service/pkg/product/repository"

	"github.com/labstack/echo/v4"
)

type productUC struct {
	repo repository.ProductRepo
}

type ProductUC interface {
	Create(ctx echo.Context, product *models.Product) (p *models.Product, err error)
}

func New(repo repository.ProductRepo) ProductUC {
	return &productUC{
		repo: repo,
	}
}

func (m *productUC) Create(ctx echo.Context, product *models.Product) (p *models.Product, err error) {
	return m.repo.Create(ctx, product)
}
