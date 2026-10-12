package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
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

func main() {
	http.HandleFunc("/", homeHandler)

	fmt.Println("Server is running at http://localhost:8080")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Server error:", err)
	}
}
