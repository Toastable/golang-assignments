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
	httpServer.HandleFunc("/create", web_server.CreateTodoHandler)
	httpServer.HandleFunc("/edit", web_server.EditTodoHandler)
	httpServer.HandleFunc("/delete", web_server.DeleteTodoHandler)
	httpServer.HandleFunc("/server-status", web_server.CheckServerStatusHandler)
	httpServer.HandleFunc("/error", web_server.ErrorHandler)

	// This works on Linux/Mac but does not on Windows due to a bug where the Registry values for file associations gets overriden to be text/plain due to a bad windows update
	// Solution is to go into the registry and manually change them back or to override golang's internal mime type map, not going to do either here just for css styling
	// fs := http.FileServer(http.Dir("/Todo/static/"))
	// httpServer.Handle("/Todo/static/", http.StripPrefix("/Todo/static", fs))

	http.ListenAndServe(address, httpServer)
}