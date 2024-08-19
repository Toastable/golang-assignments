package web_server

import (
	"bytes"
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

type PostRequestBody struct {
	Text   string 
	Status bool
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

	editTemplate := template.Must(template.ParseFiles("templates/new.html"))

	editTemplate.Execute(wr, nil)
}

func CreateTodoHandler(wr http.ResponseWriter, req *http.Request) {

	okChannel := make(chan string)
	errorChannel := make(chan int)

	go func() {
		defer close(okChannel)
		defer close(errorChannel)

		postData := PostRequestBody {
			Text: req.FormValue("todo-text"),
			Status: false,
		}

		fmt.Println(postData)

		jsonData, err := json.Marshal(postData)
		if err != nil {
			fmt.Println(err)
			errorChannel <- http.StatusInternalServerError
			return
		}

		resp, getError := http.Post(apiBaseAddress, 
			"application/json",
			bytes.NewBuffer(jsonData),
		)

		if getError != nil {
			fmt.Println(getError)
			errorChannel <- http.StatusInternalServerError
			return
		}

		defer resp.Body.Close()

		okChannel <- ""
	}()

	select {
	case <-okChannel:
		http.Redirect(wr, req, homeAddress, http.StatusFound)
	case <-errorChannel:
		http.Redirect(wr, req, errorAddress, http.StatusFound)
	}
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