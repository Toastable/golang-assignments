package web_server

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"strings"
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

type editPageViewModel struct {
	Todo todo_service.Todo 
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
	id := strings.TrimPrefix(req.URL.Path, "/edit/")

	getAddress := fmt.Sprintf("%s/%s", apiBaseAddress, id)
	resp, getError := http.Get(getAddress)

	if getError != nil {
		http.Redirect(wr, req, errorAddress, http.StatusFound)
		return
	}

	defer resp.Body.Close()

	responseBody, ioErr := io.ReadAll(resp.Body)

	if ioErr != nil {
		http.Redirect(wr, req, errorAddress, http.StatusFound)
		return
	}

	var todo todo_service.Todo
	jsonErr := json.Unmarshal(responseBody, &todo)

	if jsonErr != nil {
		http.Redirect(wr, req, errorAddress, http.StatusFound)
		return
	}

	viewModel := editPageViewModel{
		Todo: todo,
	}

	editTemplate := template.Must(template.ParseFiles("templates/edit.html"))
	editTemplate.Execute(wr, viewModel)
}

func DeleteTodoHandler(wr http.ResponseWriter, req *http.Request) {
	okChannel := make(chan string)
	errorChannel := make(chan int)

	go func() {
		defer close(okChannel)
		defer close(errorChannel)

		id := strings.TrimPrefix(req.URL.Path, "/delete/")

		deleteAddress := fmt.Sprintf("%s/%s", apiBaseAddress, id)
		deleteRequest, deleteErr := http.NewRequest(http.MethodDelete, deleteAddress, nil)

		client := &http.Client{}
		_, deleteRequestError := client.Do(deleteRequest)

		if deleteRequestError != nil {
			fmt.Println(deleteErr)
			errorChannel <- http.StatusInternalServerError
			return
		}
		okChannel <- ""
	}()

	select {
	case <-okChannel:
		http.Redirect(wr, req, homeAddress, http.StatusFound)
	case <-errorChannel:
		http.Redirect(wr, req, errorAddress, http.StatusFound)
	}
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