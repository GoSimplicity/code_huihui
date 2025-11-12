package di

import (
	"net/http"

	"github.com/GoSimplicity/code_huihui/internal/api"
	"github.com/gin-gonic/gin"
	"github.com/samber/do"
)

func InitGinServer(i *do.Injector) *gin.Engine {
	server := gin.Default()
	userHdl := do.MustInvoke[*api.UserHandler](i)
	userHdl.RegisterRoutes(server)
	server.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})

	server.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Welcome to Code Huihui API",
			"version": "1.0.0",
		})
	})

	return server
}
