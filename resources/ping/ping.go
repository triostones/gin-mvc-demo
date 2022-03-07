package ping

import (
	"github.com/triostones/gin-mvc-demo/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PingResource struct {
	path string
}

func (resource *PingResource) get(rg *gin.RouterGroup) {
	rg.GET(resource.path,
		func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"status":  http.StatusOK,
				"success": true,
				"message": "",
				"item":    (&services.PingService{}).Ping(),
			})
		},
	)
}

func (resource *PingResource) post(rg *gin.RouterGroup) {
	rg.POST(resource.path,
		func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"status":  http.StatusOK,
				"success": true,
				"message": "",
				"item":    (&services.PingService{}).Ping(),
			})
		},
	)
}
