package main

import (
	"go-blog/database"
	"go-blog/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	database.Connect()

	r.LoadHTMLGlob("templates/*")
	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "layout.html", nil)
	})

	routes.UserRoutes(r)
	routes.PostRoutes(r)

	r.Run(":8080")
}
