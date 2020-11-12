package db

import (
	"context"
	"strconv"
	"todo/api/models"
)

func GetAllTodos() (models.TodoSlice, error) {
	return models.Todos().All(context.Background(), DB)
}

func GetOneTodo(n string) ( *models.Todo, error) {
	id, err := strconv.Atoi(n)
	if err != nil {
		return nil, err
	}
	return models.FindTodo(context.Background(), DB, id)
}
