package appserver

import (
	"github.com/Mrityunjoy99/sample-go/src/application"
	"github.com/Mrityunjoy99/sample-go/src/common/config"
	"github.com/Mrityunjoy99/sample-go/src/domain/service"
	"github.com/Mrityunjoy99/sample-go/src/infrastructure/database"
	"github.com/Mrityunjoy99/sample-go/src/repository"
	"github.com/gin-gonic/gin"
)

func Start() {
	c, err := config.NewConfig()
	if err != nil {
		panic(err)
	}

	db, err := database.Connect(c)
	if err != nil {
		panic(err)
	}

	r := repository.NewRepository(db)
	domainService := service.NewServiceRegistry(c)
	s := application.NewService(r, domainService)

	g := gin.Default()
	RegisterRoutes(g, *s, *domainService)

	err = g.Run()
	if err != nil {
		panic(err)
	}
}
