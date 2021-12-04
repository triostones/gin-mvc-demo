package ping

import "github.com/gin-gonic/gin"

func RegisterPingGroup(engine *gin.Engine) {
	pingGroup := engine.Group("/ping")
	registerPing(pingGroup)
}
