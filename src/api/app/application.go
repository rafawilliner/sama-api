package app

import (
	"github.com/gin-gonic/gin"
	"github.com/rafawilliner/sama-api/src/api/config/database"
	"github.com/rafawilliner/sama-api/src/api/infrastructure/dependencies"
)

const (
	port = ":8080"
)

func Start() {

	dependencies := dependencies.StartConnection{StoreConnection: new(database.GormConnection)}
	handlers := dependencies.Start()
	router := gin.Default()
	group := router.Group("/sama")
	group.POST("pet", handlers.PetCreate.Handle)

	err := router.Run(port)
	if err != nil {
		panic(err)
	}
}
