package http

import (
	"net/http"
	"v-commerce-service/models"
	"v-commerce-service/pkg/product/usecase"

	"github.com/labstack/echo/v4"
)

type producthandler struct {
	uc usecase.ProductUC
}

func Init(uc usecase.ProductUC, g *echo.Group) {
	var (
		handler = &producthandler{
			uc: uc,
		}

		productRoutes = g.Group("/products")
	)

	productRoutes.POST("", handler.create)
}

func (h *producthandler) create(ctx echo.Context) (err error) {
	var (
		product = &models.Product{}
	)

	if err = ctx.Bind(product); err != nil {
		return
	}

	if product, err = h.uc.Create(ctx, product); err != nil {
		return
	}

	response := make(map[string]interface{})
	response["product"] = product

	return ctx.JSON(http.StatusOK, response)
}
