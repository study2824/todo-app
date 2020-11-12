package db

import (
	"context"
	"github.com/volatiletech/sqlboiler/boil"
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

func AddTodo(reqTodo models.Todo) error {
	return reqTodo.Insert(context.Background(), DB, boil.Infer())
}
