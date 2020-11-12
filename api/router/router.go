package router

import (
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

	r.GET("/todo", ctrl.GetAllTodos)
	r.GET("/todo/:id", ctrl.GetOneTodo)

	return r
}
