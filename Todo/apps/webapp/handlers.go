package main

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
	aliveAddress = "http://localhost:3000/api/alive"
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

type statusPageViewModel struct {
	Status bool
}

type PostRequestBody struct {
	Text   string 
	Status bool
}

type PatchRequestBody struct {
	ID     string
	Text   string
	Status bool
}

func HomepageHandler(wr http.ResponseWriter, req *http.Request) {
	viewModel := homepageViewModel{
		Todos: make([]todo_service.Todo, 0),
	}

	resp, getError := http.Get(apiBaseAddress)

	if getError != nil {
		http.Redirect(wr, req, errorAddress, http.StatusFound)
		return
	}

	if resp.StatusCode != http.StatusOK {
		http.Redirect(wr, req, errorAddress, http.StatusFound)
		return
	}

	defer resp.Body.Close()

	responseBody, ioErr := io.ReadAll(resp.Body)

	if ioErr != nil {
		http.Redirect(wr, req, errorAddress, http.StatusFound)
		return
	}

	todos := make([]todo_service.Todo, 0)
	jsonErr := json.Unmarshal(responseBody, &todos)

	if jsonErr != nil {
		http.Redirect(wr, req, errorAddress, http.StatusFound)
		return
	}

	viewModel.Todos = todos
	homepageTemplate := template.Must(template.ParseFiles("templates/homepage.html"))
	homepageTemplate.Execute(wr, viewModel)
}

func NewTodoHandler(wr http.ResponseWriter, req *http.Request) {
	editTemplate := template.Must(template.ParseFiles("templates/new.html"))

	editTemplate.Execute(wr, nil)
}

func CreateTodoHandler(wr http.ResponseWriter, req *http.Request) {
	postData := PostRequestBody {
		Text: req.FormValue("todo-text"),
		Status: false,
	}

	jsonData, err := json.Marshal(postData)
	if err != nil {
		http.Redirect(wr, req, errorAddress, http.StatusFound)
		return
	}

	resp, getError := http.Post(apiBaseAddress, 
		"application/json",
		bytes.NewBuffer(jsonData),
	)

	if getError != nil {
		http.Redirect(wr, req, errorAddress, http.StatusFound)
		return
	}
	
	if resp.StatusCode != http.StatusOK {
		http.Redirect(wr, req, errorAddress, http.StatusFound)
		return
	}

	defer resp.Body.Close()

	http.Redirect(wr, req, homeAddress, http.StatusFound)
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

func UpdateHandler(wr http.ResponseWriter, req *http.Request) {
	patchData := PatchRequestBody {
		ID:  req.FormValue("todo-id"),
		Text: req.FormValue("todo-text"),
		Status: req.FormValue("todo-status") == "on",
	}

	jsonData, err := json.Marshal(patchData)
	if err != nil {
		http.Redirect(wr, req, errorAddress, http.StatusFound)
		return
	}

	patchRequest, patchErr := http.NewRequest(http.MethodPatch, apiBaseAddress, bytes.NewBuffer(jsonData))

	if patchErr != nil {
		http.Redirect(wr, req, errorAddress, http.StatusFound)
		return
	}

	client := &http.Client{}
	resp, patchRequestError := client.Do(patchRequest)

	if resp.StatusCode != http.StatusOK {
		http.Redirect(wr, req, errorAddress, http.StatusFound)
		return
	}

	if patchRequestError != nil {
		http.Redirect(wr, req, errorAddress, http.StatusFound)
		return
	}

	defer resp.Body.Close()

	http.Redirect(wr, req, homeAddress, http.StatusFound)
}

func DeleteTodoHandler(wr http.ResponseWriter, req *http.Request) {

	id := strings.TrimPrefix(req.URL.Path, "/delete/")

	deleteAddress := fmt.Sprintf("%s/%s", apiBaseAddress, id)
	deleteRequest, deleteErr := http.NewRequest(http.MethodDelete, deleteAddress, nil)

	if deleteErr != nil {
		http.Redirect(wr, req, errorAddress, http.StatusFound)
		return
	}
	

	client := &http.Client{}
	resp, deleteRequestError := client.Do(deleteRequest)

	if deleteRequestError != nil {
		http.Redirect(wr, req, errorAddress, http.StatusFound)
		return
	}

	if resp.StatusCode != http.StatusOK {
		http.Redirect(wr, req, errorAddress, http.StatusFound)
		return
	}

	defer resp.Body.Close()

	http.Redirect(wr, req, homeAddress, http.StatusFound)
}

func CheckServerStatusHandler(wr http.ResponseWriter, req *http.Request) {
	resp, _ := http.Get(aliveAddress)

	viewModel := statusPageViewModel{
		Status: resp.StatusCode == http.StatusOK,
	}

	statusTemplate := template.Must(template.ParseFiles("templates/server-status.html"))
	statusTemplate.Execute(wr, viewModel)
}

func ErrorHandler(wr http.ResponseWriter, req *http.Request) {
	editTemplate := template.Must(template.ParseFiles("templates/error.html"))

	editTemplate.Execute(wr, nil)
}