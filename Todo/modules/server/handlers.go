package server

import (
	"context"
	"encoding/json"
	"net/http"
	"todo_inmemory_service"
)

const defaultTimeout = 30

type PostRequestBody struct {
	Text   string
	Status bool
}

type PatchRequestBody struct {
	ID     string
	Text   string
	Status bool
}

func UpdateHandler(service *todo_inmemory_service.TodoService) http.HandlerFunc {
	return func(wr http.ResponseWriter, req *http.Request) {
		context, cancelContext := context.WithTimeout(req.Context(), defaultTimeout)
		defer cancelContext()

		okChannel := make(chan []byte)
		errorChannel := make(chan int)

		requestBody := unmarshalRequestBody[PatchRequestBody](wr, req)

		var encodedJson []byte
		var errorResponseCode int

		service.Update(requestBody.ID, requestBody.Text, requestBody.Status)

		select {
		case <-context.Done():
			cancelContext()
			wr.WriteHeader(http.StatusRequestTimeout)
		case encodedJson = <-okChannel:
			wr.WriteHeader(http.StatusOK)
			wr.Header().Set("Content-Type", "application/json")
			wr.Write(encodedJson)
		case errorResponseCode = <-errorChannel:
			wr.WriteHeader(errorResponseCode)
		}
	}
}

func CreateHandler(service *todo_inmemory_service.TodoService) http.HandlerFunc {
	return func(wr http.ResponseWriter, req *http.Request) {
		context, cancelContext := context.WithTimeout(req.Context(), defaultTimeout)
		defer cancelContext()

		okChannel := make(chan []byte)
		errorChannel := make(chan int)

		requestBody := unmarshalRequestBody[PostRequestBody](wr, req)

		go createTodo(requestBody, service, &okChannel, &errorChannel)

		var encodedJson []byte
		var errorResponseCode int

		select {
		case <-context.Done():
			cancelContext()
			wr.WriteHeader(http.StatusRequestTimeout)
		case encodedJson = <-okChannel:
			wr.WriteHeader(http.StatusOK)
			wr.Header().Set("Content-Type", "application/json")
			wr.Write(encodedJson)
		case errorResponseCode = <-errorChannel:
			wr.WriteHeader(errorResponseCode)
		}
	}
}

func GetAllHandler(service *todo_inmemory_service.TodoService) http.HandlerFunc {
	return func(wr http.ResponseWriter, req *http.Request) {

		context, cancelContext := context.WithTimeout(req.Context(), defaultTimeout)
		defer cancelContext()

		okChannel := make(chan []byte)
		errorChannel := make(chan int)

		go retrieveTodos(service, &okChannel, &errorChannel)

		var encodedJson []byte
		var errorResponseCode int

		select {
		case <-context.Done():
			cancelContext()
			wr.WriteHeader(http.StatusRequestTimeout)
		case encodedJson = <-okChannel:
			wr.WriteHeader(http.StatusOK)
			wr.Header().Set("Content-Type", "application/json")
			wr.Write(encodedJson)
		case errorResponseCode = <-errorChannel:
			wr.WriteHeader(errorResponseCode)
		}
	}
}

func unmarshalRequestBody[T comparable](wr http.ResponseWriter, req *http.Request) T {
	var requestBody T
	err := json.NewDecoder(req.Body).Decode(&requestBody)

	if err != nil {
		wr.WriteHeader(http.StatusInternalServerError)
	}

	return requestBody
}

func createTodo(body PostRequestBody, service *todo_inmemory_service.TodoService, okChan *chan []byte, errorChan *chan int) {
	defer close(*okChan)
	defer close(*errorChan)

	if body.Text == "" {
		*errorChan <- http.StatusInternalServerError
		return
	}

	id, err := service.Create(body.Text, body.Status)

	if err != nil {
		*errorChan <- http.StatusInternalServerError
		return
	}

	encodedResponse, err := json.Marshal(id)

	if err != nil {
		*errorChan <- http.StatusInternalServerError
	}

	*okChan <- encodedResponse
}

func retrieveTodos(service *todo_inmemory_service.TodoService, okChan *chan []byte, errorChan *chan int) {
	todos, err := service.GetAll()

	defer close(*okChan)
	defer close(*errorChan)

	if err != nil {
		*errorChan <- http.StatusInternalServerError
		return
	}
	encodedTodos, err := json.Marshal(todos)

	if err != nil {
		*errorChan <- http.StatusInternalServerError
		return
	}

	*okChan <- encodedTodos
}
