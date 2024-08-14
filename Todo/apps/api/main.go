package main

import (
	"net/http"
	"server"
	"todo_inmemory_service"
)

const address = ":3000"

func main() {
	httpServer := http.NewServeMux()
	todoService := todo_inmemory_service.TodoService{}

	setupRouting(httpServer, &todoService)

	http.ListenAndServe(address, httpServer)
}

func setupRouting(httpServer *http.ServeMux, service *todo_inmemory_service.TodoService) {
	httpServer.Handle("GET /api/todo/", http.HandlerFunc(server.GetAllHandler(service)))
}
