package main

import (
	"net/http"
	"web_server"
)

const address = ":3001"

// var builtinMimeTypesLower = map[string]string{
// 	".css":  "text/css; charset=utf-8",
// 	".gif":  "image/gif",
// 	".htm":  "text/html; charset=utf-8",
// 	".html": "text/html; charset=utf-8",
// 	".jpg":  "image/jpeg",
// 	".js":   "application/javascript",
// 	".wasm": "application/wasm",
// 	".pdf":  "application/pdf",
// 	".png":  "image/png",
// 	".svg":  "image/svg+xml",
// 	".xml":  "text/xml; charset=utf-8",
// }

// func staticFileGetMimeType(ext string) string {
// 	if v, ok := builtinMimeTypesLower[ext]; ok {
// 		return v
// 	}
// 	return mime.TypeByExtension(ext)
// }

func main() {
	httpServer := http.NewServeMux()
	httpServer.HandleFunc("/", web_server.HomepageHandler)
	httpServer.HandleFunc("/new", web_server.NewTodoHandler)
	httpServer.HandleFunc("/edit", web_server.EditTodoHandler)
	httpServer.HandleFunc("/delete", web_server.DeleteTodoHandler)
	httpServer.HandleFunc("/server-status", web_server.CheckServerStatusHandler)
	httpServer.HandleFunc("/error", web_server.ErrorHandler)


	// fs := http.FileServer(http.Dir("/Todo/static/"))
	// httpServer.Handle("/Todo/static/", http.StripPrefix("/Todo/static", fs))

	http.ListenAndServe(address, httpServer)
}