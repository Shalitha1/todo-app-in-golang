package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

type Todo struct {
	ID        int
	Title     string
	Completed bool
}

var todos = []Todo{
	{
		ID:        1,
		Title:     "Learn go",
		Completed: false,
	},
	{
		ID:        2,
		Title:     "Build todo app",
		Completed: false,
	},
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, "Could not load webpage", http.StatusInternalServerError)
		log.Print(err)
		return
	}
	err = tmpl.Execute(w, todos)
	if err != nil {
		log.Println(err)
	}
}

func createTodoHandler(w http.ResponseWriter, r *http.Request) {
	title := r.FormValue("title")

	newTodo := Todo{
		ID:        len(todos) + 1,
		Title:     title,
		Completed: false,
	}

	todos = append(todos, newTodo)

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func deleteTodoHandler(w http.ResponseWriter, r *http.Request) {
	idString := r.FormValue("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		http.Error(w, "Invalid Todo ID", http.StatusBadRequest)
		return
	}
	for i, todo := range todos {
		if todo.ID == id {

		}
	}
}

func main() {
	http.HandleFunc("/", homeHandler)

	fmt.Println("Server is running at http://localhost:8080")

	http.HandleFunc("/todos/create", createTodoHandler)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Server error:", err)
	}
}
