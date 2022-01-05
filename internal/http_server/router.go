package http_server

import (
	"github.com/gin-gonic/gin"
)

func helloworld(c *gin.Context)  {
	c.JSON(200, gin.H{
		"message": "hello world!",
	})
}

func installController(g *gin.Engine) *gin.Engine {
	v1 := g.Group("/v1")
	{
		v1.GET("/hello", helloworld)
	}
	return g
}
