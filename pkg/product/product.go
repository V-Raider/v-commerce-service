package product

import (
	"v-commerce-service/pkg/product/repository"
	"v-commerce-service/pkg/product/transport/http"
	"v-commerce-service/pkg/product/usecase"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func Init(db *gorm.DB, g *echo.Group) {
	repo := repository.New(db)
	uc := usecase.New(repo)
	http.Init(uc, g)
}
