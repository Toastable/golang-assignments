package server

import (
	"context"
	"encoding/json"
	"net/http"
	"todo_inmemory_service"
)

const defaultTimeout = 30

func GetAllHandler(service *todo_inmemory_service.TodoService) http.HandlerFunc {
	return func(wr http.ResponseWriter, req *http.Request) {

		context, cancelContext := context.WithTimeout(req.Context(), defaultTimeout)
		defer cancelContext()

		okChannel := make(chan []byte)
		errorChannel := make(chan int)

		go retrieveTodosFromService(service, &okChannel, &errorChannel)

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

func retrieveTodosFromService(service *todo_inmemory_service.TodoService, okChan *chan []byte, errorChan *chan int) {
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
