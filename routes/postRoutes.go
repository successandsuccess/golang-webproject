package routes

import (
	"go-blog/controllers"
	"go-blog/middleware"

	"github.com/gin-gonic/gin"
)

func PostRoutes(r *gin.Engine) {
	auth := r.Group("/posts")
	auth.Use(middleware.AuthMiddleware())

	auth.POST("/", controllers.CreatePost)
	auth.GET("/", controllers.GetPosts)
	auth.PUT("/:id", controllers.UpdatePost)
	auth.DELETE("/:id", controllers.DeletePost)
}
