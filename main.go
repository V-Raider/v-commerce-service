package main

import (
	"v-commerce-service/service"

	"github.com/sirupsen/logrus"
)

func main() {
	if err := service.Run(); err != nil {
		logrus.Fatalln(err)
	}
}
