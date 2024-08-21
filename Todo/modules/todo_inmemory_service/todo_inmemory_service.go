package todo_inmemory_service

import (
	"errors"
	"sort"
	"strings"
	"sync"
	"todo_service"

	"github.com/google/uuid"
)

var errGetTodoNotFoundError = errors.New("could not find todo with that ID")
var errGenerateUuidError = errors.New("could not generate uuid")

type TodoService struct {
	todos []todo_service.Todo
	mutex sync.Mutex
}

func (t *TodoService) Create(text string, status bool, id string) (string, error) {
	t.mutex.Lock()
	defer t.mutex.Unlock()

	if id == "" {
		newId, err := uuid.NewV7()

		if err != nil {
			return "", errGenerateUuidError
		}

		id = newId.String()
	}

	newTodo := todo_service.Todo{
		ID:     id,
		Status: status,
		Text:   text,
	}

	t.todos = append(t.todos, newTodo)

	return newTodo.ID, nil
}

func (t *TodoService) Get(id string) (todo_service.Todo, error) {
	t.mutex.Lock()
	index, err := t.findIndexByID(id)

	defer t.mutex.Unlock()

	if err != nil {
		return todo_service.Todo{}, err
	}

	return t.todos[index], nil
}

func (t *TodoService) GetAll() ([]todo_service.Todo, error) {
	t.mutex.Lock()
	defer t.mutex.Unlock()
	return t.todos, nil
}

func (t *TodoService) Update(id string, text string, status bool) (string, error) {
	t.mutex.Lock()
	defer t.mutex.Unlock()

	index, err := t.findIndexByID(id)

	if err != nil {
		return "", err
	}

	if len(text) > 0 {
		t.todos[index].Text = text
	}

	t.todos[index].Status = status

	return id, nil
}

func (t *TodoService) Delete(id string) error {
	t.mutex.Lock()
	defer t.mutex.Unlock()
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
