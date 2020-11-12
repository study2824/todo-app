package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"todo/api/db"
)

func (Controller) GetAllTodo(c *gin.Context) {
	tasks, err := db.GetAllTodo()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"tasks": tasks})
}
