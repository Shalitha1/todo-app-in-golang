package main

import (
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"
)

// Home handler

func homeHandler(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query(
		"SELECT id, title, completed FROM todos ORDER BY id",
	)
	if err != nil {
		http.Error(w, "Could not load todos", http.StatusInternalServerError)
		log.Println("Query error:", err)
		return
	}
	defer rows.Close()

	var todos []Todo

	for rows.Next() {
		var todo Todo

		err := rows.Scan(&todo.ID, &todo.Title, &todo.Completed)
		if err != nil {
			http.Error(w, "Could not read todos", http.StatusInternalServerError)
			log.Println("Scan error:", err)
			return
		}

		todos = append(todos, todo)
	}

	if err := rows.Err(); err != nil {
		http.Error(w, "Could not read todos", http.StatusInternalServerError)
		log.Println("Rows error:", err)
		return
	}

	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, "Could not load webpage", http.StatusInternalServerError)
		log.Println(err)
		return
	}

	if err := tmpl.Execute(w, todos); err != nil {
		log.Println("Template error:", err)
	}
}

// Create
func createTodoHandler(w http.ResponseWriter, r *http.Request) {
	title := r.FormValue("title")
	title = strings.TrimSpace(title)
	if title == "" {
		http.Error(w, "Title is empty", http.StatusBadRequest)
		return
	}

	newTodo := Todo{
		ID:        nextTodoId,
		Title:     title,
		Completed: false,
	}

	todos = append(todos, newTodo)
	nextTodoId++

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
