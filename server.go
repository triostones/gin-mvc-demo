package main

import (
	"github.com/gin-gonic/gin"
	"github.com/triostones/triostones-backend-gin/src/config"
	"github.com/triostones/triostones-backend-gin/src/entity"
	"github.com/triostones/triostones-backend-gin/src/resources"
)

func main() {

	config.Initialize()

	entity.AutoMigrate()

	engine := gin.Default()

	resources.RegisterResources(engine)

	engine.Run()

}
