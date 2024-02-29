package appserver

import (
	"github.com/Mrityunjoy99/sample-go/src/application"
	"github.com/Mrityunjoy99/sample-go/src/common/config"
	"github.com/Mrityunjoy99/sample-go/src/infrastructure/database"
	"github.com/Mrityunjoy99/sample-go/src/infrastructure/repository"
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
	s:=application.NewService(*r)

	g := gin.Default()
	RegisterRoutes(g,*s)

	g.Run()
}
