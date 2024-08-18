package web_server

import (
	"html/template"
	"net/http"
	"todo_service"
)

const defaultTimeout = 60

type homepageViewModel struct {
	Todos []todo_service.Todo
}


func HomepageHandler(wr http.ResponseWriter, req *http.Request) {
	viewModel := homepageViewModel{
		Todos: make([]todo_service.Todo, 0),
	}

	homepageTemplate := template.Must(template.ParseFiles("templates/homepage.html"))

	homepageTemplate.Execute(wr, viewModel)

}

func NewTodoHandler(wr http.ResponseWriter, req *http.Request) {
	viewModel := homepageViewModel{
		Todos: make([]todo_service.Todo, 0),
	}

	newTemplate := template.Must(template.ParseFiles("templates/new.html"))

	newTemplate.Execute(wr, viewModel)
}

func EditTodoHandler(wr http.ResponseWriter, req *http.Request) {
	viewModel := homepageViewModel{
		Todos: make([]todo_service.Todo, 0),
	}

	editTemplate := template.Must(template.ParseFiles("templates/edit.html"))

	editTemplate.Execute(wr, viewModel)
}

func CheckServerStatusHandler(wr http.ResponseWriter, req *http.Request) {
	viewModel := homepageViewModel{
		Todos: make([]todo_service.Todo, 0),
	}

	editTemplate := template.Must(template.ParseFiles("templates/status.html"))

	editTemplate.Execute(wr, viewModel)
}