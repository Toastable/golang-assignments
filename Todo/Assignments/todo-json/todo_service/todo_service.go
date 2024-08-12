package todo_service

type Todo struct {
	ID     string `json:"id"`
	Status bool   `json:"status"`
	Text   string `json:"text"`
}

type TodoService interface {
	Create(text string) error
	Get(id string) (Todo, error)
	GetAll() ([]Todo, error)
	Update(id string, text string, status bool) error
	Delete(id string) error
}
