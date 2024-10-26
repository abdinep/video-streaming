package main

import (
	"net/http"

	"github.com/abdinep/video_streaming/handler"
	"github.com/gin-gonic/gin"
)

func main() {
	c := gin.Default()
	c.LoadHTMLGlob("templates/*")
	c.GET("/", handler.Stream)
	c.GET("/video", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.html", gin.H{})
	})
	c.StaticFS("/hls", http.Dir("./Out"))
	c.Run(":8080")
}
