package main

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"
	"time"
	"todo_inmemory_service"
)

const defaultTimeout = time.Millisecond * 90

type PostRequestBody struct {
	Text   string
	Status bool
}

type PatchRequestBody struct {
	ID     string
	Text   string
	Status bool
}

func AliveHandler(wr http.ResponseWriter, req *http.Request) {
	wr.WriteHeader(http.StatusOK)
}

func DeleteHandler(service *todo_inmemory_service.TodoService) http.HandlerFunc {
	return func(wr http.ResponseWriter, req *http.Request) {
		context, cancelContext := context.WithTimeout(req.Context(), defaultTimeout)
		defer cancelContext()

		responseChannel := make(chan int)

		id := req.PathValue("id")

		var responseCode int

		go deleteTodo(id, service, &responseChannel)

		select {
		case <-context.Done():
			cancelContext()
			wr.WriteHeader(http.StatusRequestTimeout)
		case responseCode = <-responseChannel:
			wr.WriteHeader(responseCode)
		}
	}
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

		go updateTodo(requestBody, service, &okChannel, &errorChannel)

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

func GetHandler(service *todo_inmemory_service.TodoService) http.HandlerFunc {
	return func(wr http.ResponseWriter, req *http.Request) {

		context, cancelContext := context.WithTimeout(req.Context(), defaultTimeout)
		defer cancelContext()

		okChannel := make(chan []byte)
		errorChannel := make(chan int)

		id := strings.TrimPrefix(req.URL.Path, "/api/todo/")

		go retrieveTodoById(id, service, &okChannel, &errorChannel)

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

func deleteTodo(id string, service *todo_inmemory_service.TodoService, responseChan *chan int) {
	defer close(*responseChan)

	err := service.Delete(id)

	if err != nil {
		*responseChan <- http.StatusNotFound
	}

	*responseChan <- http.StatusOK
}

func updateTodo(body PatchRequestBody, service *todo_inmemory_service.TodoService, okChan *chan []byte, errorChan *chan int) {
	defer close(*okChan)
	defer close(*errorChan)

	id, err := service.Update(body.ID, body.Text, body.Status)

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

func createTodo(body PostRequestBody, service *todo_inmemory_service.TodoService, okChan *chan []byte, errorChan *chan int) {
	defer close(*okChan)
	defer close(*errorChan)

	if body.Text == "" {
		*errorChan <- http.StatusInternalServerError
		return
	}

	id, err := service.Create(body.Text, body.Status, "")

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

func retrieveTodoById(id string, service *todo_inmemory_service.TodoService, okChan *chan []byte, errorChan *chan int) {
	todo, err := service.Get(id)

	defer close(*okChan)
	defer close(*errorChan)

	if err != nil {
		*errorChan <- http.StatusInternalServerError
		return
	}
	encodedTodo, err := json.Marshal(todo)

	if err != nil {
		*errorChan <- http.StatusInternalServerError
		return
	}

	*okChan <- encodedTodo
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
