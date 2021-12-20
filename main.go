package main

import (
	"github.com/gin-gonic/gin"
	"github.com/triostones/gin-mvc-demo/src/models"
	"github.com/triostones/gin-mvc-demo/src/resources"
)

func main() {

	models.AutoMigrate()

	engine := gin.Default()

	resources.RegisterResources(engine)

	engine.Run()

}
