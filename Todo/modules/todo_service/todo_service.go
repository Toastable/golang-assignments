package todo_service

import (
	"fmt"
)

type Todo struct {
	ID     string `json:"ID"`
	Status bool   `json:"Status"`
	Text   string `json:"Text"`
}

type TodoService interface {
	Create(text string, status bool) (string, error)
	GetAll() ([]Todo, error)
	Update(id string, text string, status bool) (string, error)
	Delete(id string) error
}

func (todo Todo) String() string {
	return fmt.Sprintf("| %-33s | %-50s | %-11t |", todo.ID, todo.Text, todo.Status)
}
