package server

import (
	"context"
	"fmt"
	"net/http"
	"todo_inmemory_service"
)

func GetAllHandler(service *todo_inmemory_service.TodoService) http.HandlerFunc {
	return func(wr http.ResponseWriter, req *http.Request) {

		context, cancelContext := context.WithTimeout(req.Context(), 30)

		defer cancelContext()
		fmt.Println(context)

		//do todo stuff here
		todos, err := service.GetAll()

		fmt.Println(todos)
		fmt.Println(err)
	}
}
