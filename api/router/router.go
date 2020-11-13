package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"todo/api/controller"
)

func Init() {
	r := router()
	r.Run("localhost:8080")
}

func router() *gin.Engine {
	r := gin.Default()

	ctrl := controller.Controller{}

	// 全てのドメインのアクセス許可
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	r.Use(cors.New(config))

	r.GET("/todo", ctrl.GetAllTodos)
	r.GET("/todo/:id", ctrl.GetOneTodo)
	r.POST("/todo/add", ctrl.AddTodo)
	r.PUT("/todo/:id/update", ctrl.UpdateTodo)
	r.DELETE("/todo/:id/delete", ctrl.DeleteTodo)
	r.POST("/search", ctrl.SearchTodo)
	return r
}
