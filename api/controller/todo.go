package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"todo/api/db"
	"todo/api/models"
)

func (Controller) GetAllTodos(c *gin.Context) {
	todos, err := db.GetAllTodos()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"todos": todos})
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

func (Controller) AddTodo(c *gin.Context) {
	reqTodo := models.Todo{
		Title: c.PostForm("title"),
		Text:  c.PostForm("text"),
	}
	err := db.AddTodo(reqTodo)
	if err != nil {
		ErrMsg(err, c)
		return
	}

	createdTodo, err := db.GetLastTodo()
	if err != nil {
		ErrMsg(err, c)
	}

	c.JSON(http.StatusOK, gin.H{"added todo": createdTodo})
}

func (Controller) UpdateTodo(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		ErrMsg(err, c)
		return
	}

	reqTodo := models.Todo{
		ID:    id,
		Title: c.PostForm("title"),
		Text:  c.PostForm("text"),
	}

	err = db.UpdateTodo(reqTodo)
	if err != nil {
		ErrMsg(err, c)
		return
	}

	c.JSON(http.StatusOK, gin.H{"updated todo": reqTodo})
}

func (Controller) DeleteTodo(c *gin.Context) {
	id := c.Param("id")
	deletedTodo, err := db.GetOneTodo(id)
	if err != nil {
		ErrMsg(err, c)
		return
	}

	err = db.DeleteTodo(id)
	if err != nil {
		ErrMsg(err, c)
		return
	}

	c.JSON(http.StatusOK, gin.H{"deleted todo": deletedTodo})

}

func (Controller) SearchTodo(c *gin.Context) {
	target := c.PostForm("target")
	keyword := c.PostForm("keyword")
	todos, err := db.SearchTodo(target, keyword)
	if err != nil {
		ErrMsg(err, c)
		return
	}

	c.JSON(http.StatusOK, gin.H{"result": todos})
}
