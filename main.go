package main

import (
	"io"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	config := ""

	r := gin.Default()

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"*"}
	r.Use(cors.New(corsConfig))

	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	r.GET("/config", func(c *gin.Context) {
		if config == "" {
			c.Status(404)
			return
		}
		c.String(200, config)
	})

	r.POST("/config", func(c *gin.Context) {
		config = c.PostForm("config")
		if config == "" {
			c.Status(404)
			return
		}
		c.Status(204)
	})

	r.GET("/proxy", func(c *gin.Context) {
		url := c.Query("url")
		if url == "" {
			c.Status(404)
			return
		}
		resp, err := http.Get(url)
		if err != nil {
			c.String(500, err.Error())
			return
		}
		defer resp.Body.Close()
		io.Copy(c.Writer, resp.Body)
	})

	r.Run(":54637")
}
