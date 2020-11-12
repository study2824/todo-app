package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"todo/api/db"
)

func (Controller) GetAllTodos(c *gin.Context) {
	todos, err := db.GetAllTodos()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"tasks": todos})
}

func (Controller) GetOneTodo(c *gin.Context) {
	id := c.Param("id")
	todo, err := db.GetOneTodo(id)
	if err != nil {
		ErrMsg(err, c)
		return
	}

	c.JSON(http.StatusOK, gin.H{"todo": todo})
}
