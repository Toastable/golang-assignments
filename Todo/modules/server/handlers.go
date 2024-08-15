package server

import (
	"context"
	"encoding/json"
	"net/http"
	"todo_inmemory_service"
)

const defaultTimeout = 30

type RequestBody struct {
	Text   string
	Status bool
}

func CreateHandler(service *todo_inmemory_service.TodoService) http.HandlerFunc {
	return func(wr http.ResponseWriter, req *http.Request) {
		context, cancelContext := context.WithTimeout(req.Context(), defaultTimeout)
		defer cancelContext()

		okChannel := make(chan []byte)
		errorChannel := make(chan int)

		requestBody := unmarshalRequestBody(wr, req)

		go createTodo(requestBody, service, &okChannel, &errorChannel)

		var encodedJson []byte
		var errorResponseCode int

		select {
		case <-context.Done():
			cancelContext()
			wr.WriteHeader(http.StatusRequestTimeout)
		case encodedJson = <-okChannel:
			wr.WriteHeader(http.StatusOK)
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
			wr.Write(encodedJson)
		case errorResponseCode = <-errorChannel:
			wr.WriteHeader(errorResponseCode)
		}
	}
}

func unmarshalRequestBody(wr http.ResponseWriter, req *http.Request) RequestBody {
	var requestBody RequestBody
	err := json.NewDecoder(req.Body).Decode(&requestBody)

	if err != nil {
		wr.WriteHeader(http.StatusInternalServerError)
	}

	return requestBody
}

func createTodo(body RequestBody, service *todo_inmemory_service.TodoService, okChan *chan []byte, errorChan *chan int) {
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
