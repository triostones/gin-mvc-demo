package main

import (
	"github.com/gin-gonic/gin"
	"github.com/triostones/triostones-backend-gin/src/models"
	"github.com/triostones/triostones-backend-gin/src/resources"
)

func main() {

	models.AutoMigrate()

	engine := gin.Default()

	resources.RegisterResources(engine)

	engine.Run()

}
