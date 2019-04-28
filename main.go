package main

import (
	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/", handler)
	return r
}


func handler(c *gin.Context) {
	c.JSON(200, gin.H{"msg": "Hello World!\n"})
}

func main() {
	r := setupRouter()
	r.Run()
}
