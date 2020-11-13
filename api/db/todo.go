package db

import (
	"context"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries/qm"
	"strconv"
	"todo/api/models"
)

func GetAllTodos() (models.TodoSlice, error) {
	return models.Todos().All(context.Background(), DB)
}

func GetOneTodo(n string) (*models.Todo, error) {
	id, err := strconv.Atoi(n)
	if err != nil {
		return nil, err
	}
	return models.FindTodo(context.Background(), DB, id)
}

func GetLastTodo() (*models.Todo, error) {
	allTodo, err := GetAllTodos()
	if err != nil {
		return nil, err
	}

	l := strconv.Itoa(len(allTodo))
	return GetOneTodo(l)
}

func AddTodo(reqTodo models.Todo) error {
	return reqTodo.Insert(context.Background(), DB, boil.Infer())
}

func UpdateTodo(reqTodo models.Todo) error {
	_, err := reqTodo.Update(context.Background(), DB, boil.Infer())
	return err
}

func DeleteTodo(id string) error {
	todo, err := GetOneTodo(id)
	if err != nil {
		return err
	}

	_, err = todo.Delete(context.Background(), DB)
	return err
}

func SearchTodo(target string,keyword string) (models.TodoSlice, error) {
	clause := target + " like ?"
	args := `%`+keyword+`%`
	return models.Todos(qm.Select("*"),
		qm.Where(clause, args)).
		All(context.Background(), DB)
}
