package main

import (
	"github.com/gin-gonic/gin"
	"github.com/triostones/gin-mvc-demo/models"
	"github.com/triostones/gin-mvc-demo/resources"
)

func main() {

	models.AutoMigrate()

	engine := gin.Default()

	resources.RegisterResources(engine)

	engine.Run()

}
