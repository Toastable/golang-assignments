package api_handlers_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"
)

const (
	apiBaseAddress = "http://localhost:3000/api/todo"
)

type PostRequestBody struct {
	Text   string
	Status bool
}

func TestGetAllHandler(t *testing.T) {
	t.Run("returns 200 response", func(t *testing.T) {
		resp, _ := http.Get(apiBaseAddress)

		assertAreEqual(t, resp.StatusCode, http.StatusOK)
	})
}

func TestCreateHandler(t *testing.T) {
	t.Run("returns 200 response with a valid post body", func(t *testing.T) {

		postData := PostRequestBody {
			Text: "An Integration Test Todo",
			Status: true,
		}

		jsonData, _ := json.Marshal(postData)
		resp, getError := http.Post(apiBaseAddress, 
			"application/json",
			bytes.NewBuffer(jsonData),
		)

		assertAreEqual(t, getError, nil)
		assertAreEqual(t, resp.StatusCode, http.StatusOK)
	})

	t.Run("returns 500 response with a malformed post body", func(t *testing.T) {
		malformedJsonData := make([]byte, 0)
		resp, _ := http.Post(apiBaseAddress, 
			"application/json",
			bytes.NewBuffer(malformedJsonData),
		)
		
		assertAreEqual(t, resp.StatusCode, http.StatusInternalServerError)
	})

	t.Run("returns 500 response and no error with an empty text field", func(t *testing.T) {
		postData := PostRequestBody {
			Text: "",
			Status: true,
		}

		jsonData, _ := json.Marshal(postData)

		resp, _ := http.Post(apiBaseAddress, 
			"application/json",
			bytes.NewBuffer(jsonData),
		)
		
		assertAreEqual(t, resp.StatusCode, http.StatusInternalServerError)
	})
}

func TestDeleteHandler(t *testing.T) {
	t.Run("returns 200 response with a valid id", func(t *testing.T) {

		deleteAddress := fmt.Sprintf("%s/%s", apiBaseAddress, "c3e6f7d0-8b5a-4e8d-9c5b-3a2b1e0f4a7d")
		deleteRequest, _ := http.NewRequest(http.MethodDelete, deleteAddress, nil)

		client := &http.Client{}
		resp, _ := client.Do(deleteRequest)

		assertAreEqual(t, resp.StatusCode, http.StatusOK)
	})

	t.Run("returns 404 response with an id that does not exist", func(t *testing.T) {

		deleteAddress := fmt.Sprintf("%s/%s", apiBaseAddress, "Some Fake ID")
		deleteRequest, _ := http.NewRequest(http.MethodDelete, deleteAddress, nil)

		client := &http.Client{}
		resp, _ := client.Do(deleteRequest)

		assertAreEqual(t, resp.StatusCode, http.StatusNotFound)
	})
}

func TestGetHandler(t *testing.T) {
	t.Run("returns 200 response", func(t *testing.T) {
		address := fmt.Sprintf("%s/%s", apiBaseAddress, "d4c3b2a1-0e1f-2d3c-4b5a-6f7e8d9c0a")
		resp, _ := http.Get(address)
		fmt.Println(resp.StatusCode)

		assertAreEqual(t, resp.StatusCode, http.StatusOK)
	})
}

func assertAreEqual[T comparable](t testing.TB, got, want T) {
	t.Helper()

	if got != want {
		t.Errorf(`got %v want %v`, got, want)
	}
}
