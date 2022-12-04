package service

import (
	"v-commerce-service/models"
	"v-commerce-service/pkg/product"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Run() (err error) {
	var (
		db *gorm.DB
		e  *echo.Echo
	)

	logrus.Infoln("service try to running ...")

	dsn := "host=localhost user=postgres password=postgres dbname=v-commerce port=5432 sslmode=disable TimeZone=Asia/Tehran"
	if db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{}); err != nil {
		return
	}

	if err = db.AutoMigrate(&models.Product{}); err != nil {
		return
	}

	e = echo.New()
	g := e.Group("/api/v1")

	product.Init(db, g)

	e.Logger.Fatal(e.Start(":1323"))

	return
}
