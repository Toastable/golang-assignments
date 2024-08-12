package todo_service

type Todo struct {
	ID     string `json:"ID"`
	Status bool   `json:"Status"`
	Text   string `json:"Text"`
}

type TodoService interface {
	Create(text string) error
	Get(id string) (Todo, error)
	GetAll() ([]Todo, error)
	Update(id string, text string, status bool) error
	Delete(id string) error
}
