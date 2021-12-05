package ping

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/triostones/triostones-backend-gin/src/services"
)

func get(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"success": true,
		"message": "",
		"item":    (&services.PingService{}).Ping(),
	})
}

func registerPing(rg *gin.RouterGroup) {
	rg.GET("/", get)
}
