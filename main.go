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

// Home handler
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

// Create
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

// Delete
func deleteTodoHandler(w http.ResponseWriter, r *http.Request) {
	idString := r.FormValue("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		http.Error(w, "Invalid Todo ID", http.StatusBadRequest)
		return
	}
	for i, todo := range todos {
		if todo.ID == id {
			todos = append(todos[:i], todos[i+1:]...)
			break
		}
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

// Update
func updateTodoHandler(w http.ResponseWriter, r *http.Request) {
	idString := r.FormValue("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		http.Error(w, "Invalid Todo ID", http.StatusBadRequest)
		return
	}
	for i, todo := range todos {
		if todo.ID == id {
			todos[i].Completed = !todo.Completed
			break
		}
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

// Edit
func editTodoHandler(w http.ResponseWriter, r *http.Request) {
	idString := r.FormValue("id")
	id, err := strconv.Atoi(idString)

	if err != nil {
		http.Error(w, "Invalid Todo ID", http.StatusBadRequest)
		return
	}

	newTitle := r.FormValue("title")

	for i, todo := range todos {
		if todo.ID == id {
			todos[i].Title = newTitle
			break
		}
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

// Main function
func main() {
	http.HandleFunc("/", homeHandler)

	fmt.Println("Server is running at http://localhost:8080")

	http.HandleFunc("/todos/create", createTodoHandler)
	http.HandleFunc("/todos/delete", deleteTodoHandler)
	http.HandleFunc("/todos/update", updateTodoHandler)
	http.HandleFunc("/todos/edit", editTodoHandler)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Server error:", err)
	}
}
