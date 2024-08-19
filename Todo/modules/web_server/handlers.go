package web_server

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"todo_service"
)

const (
	apiBaseAddress = "http://localhost:3000/api/todo"
	errorAddress = "http://localhost:3001/error"
	homeAddress = "http://localhost:3001"
)

type homepageViewModel struct {
	Todos []todo_service.Todo
}


func HomepageHandler(wr http.ResponseWriter, req *http.Request) {
	viewModel := homepageViewModel{
		Todos: make([]todo_service.Todo, 0),
	}

	okChannel := make(chan []todo_service.Todo)
	errorChannel := make(chan int)

	go func() {
		defer close(okChannel)
		defer close(errorChannel)

		resp, getError := http.Get(apiBaseAddress)

		if getError != nil {
			fmt.Println(getError)
			errorChannel <- http.StatusInternalServerError
			return
		}

		defer resp.Body.Close()

		responseBody, ioErr := io.ReadAll(resp.Body)

		if ioErr != nil {
			errorChannel <- http.StatusInternalServerError
			return
		}

		todos := make([]todo_service.Todo, 0)
		jsonErr := json.Unmarshal(responseBody, &todos)

		if jsonErr != nil {
			errorChannel <- http.StatusInternalServerError
			return
		}

		okChannel <- todos
	}()
	
	var todos []todo_service.Todo

	select {
	case todos = <-okChannel:
		viewModel.Todos = todos
		homepageTemplate := template.Must(template.ParseFiles("templates/homepage.html"))
		homepageTemplate.Execute(wr, viewModel)
	case <-errorChannel:
		http.Redirect(wr, req, errorAddress, http.StatusFound)
	}
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

func DeleteTodoHandler(wr http.ResponseWriter, req *http.Request) {
	//Do delete stuff

	http.Redirect(wr, req, homeAddress, http.StatusOK)
}

func CheckServerStatusHandler(wr http.ResponseWriter, req *http.Request) {
	viewModel := homepageViewModel{
		Todos: make([]todo_service.Todo, 0),
	}

	statusTemplate := template.Must(template.ParseFiles("templates/status.html"))

	statusTemplate.Execute(wr, viewModel)
}

func ErrorHandler(wr http.ResponseWriter, req *http.Request) {
	viewModel := homepageViewModel{
		Todos: make([]todo_service.Todo, 0),
	}

	editTemplate := template.Must(template.ParseFiles("templates/error.html"))

	editTemplate.Execute(wr, viewModel)
}