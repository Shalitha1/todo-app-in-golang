package main

import (
	"fmt"
	"net/http"
)

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
