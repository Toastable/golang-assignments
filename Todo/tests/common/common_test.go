package main

import (
	"common"
	"reflect"
	"testing"
	"todo_inmemory_service"
	"todo_service"
)

const testFilePath string = "todo-test-data.json"

func TestPopulateInMemoryTodos(t *testing.T) {
	t.Run("populates todos from test json file", func(t *testing.T) {

		inMemService := todo_inmemory_service.TodoService{}

		want := make([]todo_service.Todo, 2)
		want = append(want, todo_service.Todo{
			ID:     "c3e6f7d0-8b5a-4e8d-9c5b-3a2b1e0f4a7d",
			Status: false,
			Text:   "Learn GoLang",
		})
		want = append(want, todo_service.Todo{
			ID:     "a1b2c3d4-5e6f-7a8b-9c0d-e1f2a3b4c5d6",
			Status: true,
			Text:   "Make Tea",
		})

		got := common.PopulateInMemoryTodos(&inMemService, testFilePath)

		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want []todo_service.Todo) {
	t.Helper()

	if reflect.DeepEqual(got, want) {
		t.Errorf("got %q want %q", got, want)
	}
}
