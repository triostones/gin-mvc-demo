package ping

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/triostones/triostones-backend-gin/src/services"
)

func get(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "success",
		"item":    (&services.PingService{}).Ping(),
	})
}

func registerPing(rg *gin.RouterGroup) {
	rg.GET("/", get)
}
