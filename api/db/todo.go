package db

import (
	"context"
	"todo/api/models"
)

func GetAllTodo() (models.TodoSlice, error) {
	return models.Todos().All(context.Background(), DB)
}
