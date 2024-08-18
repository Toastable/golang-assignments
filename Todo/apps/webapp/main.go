package main

import (
	"net/http"
	"web_server"
)

const address = ":3001"

func main() {
	httpServer := http.NewServeMux()
	httpServer.HandleFunc("/", web_server.HomepageHandler)
	httpServer.HandleFunc("/new", web_server.NewTodoHandler)
	httpServer.HandleFunc("/edit", web_server.EditTodoHandler)
	httpServer.HandleFunc("/server-status", web_server.CheckServerStatusHandler)

	http.ListenAndServe(address, httpServer)
}