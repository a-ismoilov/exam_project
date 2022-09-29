package api

import "github.com/gin-gonic/gin"

func Server(c *controllers) *gin.Engine {
	e := gin.Default()
	v1 := e.Group("api/v1")
	v1.GET("/get-words", c.ReadWords())
	v1.POST("/set-words", c.WriteWords())

	return e
}
