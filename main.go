package main

import (
	"fmt"

	"github.com/HAL-RO-Developer/test_server/model"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Static("/js", "./public/js")
	r.Static("/image", "./public/image")
	r.Static("/css", "./public/css")

	//r.LoadHTMLGlob("view/*")
	//r.NoRoute(func(c *gin.Context) {
	//	c.HTML(200, "index.html", nil)
	//})

	r.POST("/api/foo", func(c *gin.Context) {
		s, _ := c.GetRawData()
		fmt.Println(string(s))
		model.LogRep.Set(string(s))
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})
	r.GET("/api/logs", func(c *gin.Context) {
		c.JSON(200, model.LogRep.Get(10))
	})
	r.Run()
}
