package models

import (
	controller "postsapi/controllers"

	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.RouterGroup) {
	r.GET("/posts", controller.ListPosts)
	r.GET("/posts/:id", controller.FindPost)
	r.POST("/posts", controller.CreatePost)
	r.DELETE("/posts/:id", controller.DeletePost)
	r.PUT("/posts", controller.UpdatePost)
}
