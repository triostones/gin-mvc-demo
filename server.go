package main

import (
	"github.com/gin-gonic/gin"
	"github.com/triostones/triostones-backend-gin/src/resources"
)

func main() {

	engine := gin.Default()

	resources.RegisterResources(engine)

	engine.Run()

}
