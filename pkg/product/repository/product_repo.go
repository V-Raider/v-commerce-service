package repository

import (
	"v-commerce-service/models"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type productRepo struct {
	db *gorm.DB
}

type ProductRepo interface {
	Create(ctx echo.Context, product *models.Product) (p *models.Product, err error)
}

func New(db *gorm.DB) ProductRepo {
	return &productRepo{
		db: db,
	}
}

func (m *productRepo) Create(ctx echo.Context, product *models.Product) (p *models.Product, err error) {
	if err = m.db.Create(&product).Error; err != nil {
		return
	}

	return product, nil
}
