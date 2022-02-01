package routers

import (
	"blog/controller"
	"blog/logger"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetUp(mode string) *gin.Engine {
	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	r.POST("/signup", controller.SignUpHandler)
	r.POST("/login", controller.LoginHandler)

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "hello gin")
	})
	return r
}
