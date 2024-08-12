package todo_inmemory_service

import (
	"errors"
	"sort"
	"strings"
	"todo_service"

	"github.com/google/uuid"
)

var errGetTodoNotFoundError = errors.New("could not find todo with that ID")
var errGenerateUuidError = errors.New("could not generate uuid")

type TodoService struct {
	todos []todo_service.Todo
}

func (t *TodoService) Create(text string) error {

	id, err := uuid.NewV7()

	if err != nil {
		return errGenerateUuidError
	}

	newTodo := todo_service.Todo{
		ID:     id.String(),
		Status: false,
		Text:   text,
	}

	t.todos = append(t.todos, newTodo)

	return nil
}

func (t *TodoService) Get(id string) (todo_service.Todo, error) {
	index, err := t.findIndexByID(id)

	if err != nil {
		return todo_service.Todo{}, err
	}

	return t.todos[index], nil
}

func (t *TodoService) GetAll() ([]todo_service.Todo, error) {
	return t.todos, nil
}

func (t *TodoService) Update(id string, text string, status bool) error {

	index, err := t.findIndexByID(id)

	if err != nil {
		return err
	}

	t.todos[index].Text = text
	t.todos[index].Status = status

	return nil
}

func (t *TodoService) Delete(id string) error {
	index, err := t.findIndexByID(id)

	if err != nil {
		return err
	}

	t.todos = append(t.todos[:index], t.todos[index+1:]...)

	return nil
}

func (t *TodoService) findIndexByID(id string) (int, error) {
	index, found := sort.Find(len(t.todos), func(i int) int {
		return strings.Compare(id, t.todos[i].ID)
	})

	if !found {
		return -1, errGetTodoNotFoundError
	}

	return index, nil
}
