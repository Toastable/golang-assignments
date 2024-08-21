package main

import (
	"net/http"
)

const address = ":3001"

func main() {
	httpServer := http.NewServeMux()
	httpServer.HandleFunc("/", HomepageHandler)
	httpServer.HandleFunc("/new", NewTodoHandler)
	httpServer.HandleFunc("/create", CreateTodoHandler)
	httpServer.HandleFunc("/edit/{id}", EditTodoHandler)
	httpServer.HandleFunc("/update", UpdateHandler)
	httpServer.HandleFunc("/delete/{id}", DeleteTodoHandler)
	httpServer.HandleFunc("/server-status", CheckServerStatusHandler)
	httpServer.HandleFunc("/error", ErrorHandler)

	// This works on Linux/Mac but does not on Windows due to a bug where the Registry values for file associations gets overriden to be text/plain due to a bad windows update
	// Solution is to go into the registry and manually change them back or to override golang's internal mime type map, not going to do either here just for css styling
	// fs := http.FileServer(http.Dir("/Todo/static/"))
	// httpServer.Handle("/Todo/static/", http.StripPrefix("/Todo/static", fs))

	http.ListenAndServe(address, httpServer)
}