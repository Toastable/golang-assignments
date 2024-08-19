package todo_inmemory_service_test

import (
	"testing"
	"todo_inmemory_service"
	"todo_service"
)

func TestCreate(t *testing.T) {
	t.Run("correctly appends new todo to list", func(t *testing.T) {
		service := todo_inmemory_service.TodoService{}

		want := make([]todo_service.Todo, 0)
		want = append(want, todo_service.Todo{
			Status: false,
			Text:   "This is a test",
		})

		_, err := service.Create("This is also a test", true)

		if err != nil {
			t.Fatalf("Expected no errors to occur but one happened anyway: %v", err)
		}

		got, _ := service.GetAll()

		assertAreEqual(t, len(got), len(want))
		sut := got[0]

		assertAreEqual(t, sut.Text, "This is also a test")
		assertAreEqual(t, sut.Status, true)
	})
}

func TestDelete(t *testing.T) {
	t.Run("correctly removes a todo from list", func(t *testing.T) {
		service := todo_inmemory_service.TodoService{}

		want := make([]todo_service.Todo, 0)
		want = append(want, todo_service.Todo{
			Status: false,
			Text:   "Learn GoLang",
		})

		service.Create("Learn GoLang", false)
		service.Create("This should be removed", true)

		todos, _ := service.GetAll()
		err := service.Delete(todos[1].ID)

		if err != nil {
			t.Fatalf("Expected no errors to occur but one happened anyway: %v", err)
		}

		got, _ := service.GetAll()

		assertAreEqual(t, len(got), len(want))
	})
}

func TestGetAll(t *testing.T) {
	t.Run("correctly retrieves todos", func(t *testing.T) {
		service := todo_inmemory_service.TodoService{}

		want := make([]todo_service.Todo, 0)
		want = append(want, todo_service.Todo{
			Status: false,
			Text:   "Learn GoLang",
		})
		want = append(want, todo_service.Todo{
			Status: true,
			Text:   "Pass this unit test",
		})
		service.Create("Learn GoLang", false)
		service.Create("Pass this unit test", true)

		got, _ := service.GetAll()

		assertAreEqual(t, len(got), len(want))

		sut := got[0]
		assertAreEqual(t, sut.Text, want[0].Text)
		assertAreEqual(t, sut.Status, want[0].Status)

		sut = got[1]
		assertAreEqual(t, sut.Text, want[1].Text)
		assertAreEqual(t, sut.Status, want[1].Status)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("correctly updates the status of a todo", func(t *testing.T) {
		service := todo_inmemory_service.TodoService{}
		service.Create("This is a test", true)

		todos, _ := service.GetAll()

		_, err := service.Update(todos[0].ID, "", false)

		if err != nil {
			t.Fatalf("Expected no errors to occur but one happened anyway: %v", err)
		}

		got, _ := service.GetAll()
		sut := got[0]
		assertAreEqual(t, sut.Status, false)
	})
}

func assertAreEqual[T comparable](t testing.TB, got, want T) {
	t.Helper()

	if got != want {
		t.Errorf(`got %v want %v`, got, want)
	}
}
