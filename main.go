package main

import (
	"fmt"
	"log"
	"net/http"
)

// Main function
func main() {

	if err := initDB(); err != nil {
		log.Fatal("Database connection failed: ", err)
	}
	defer db.Close()

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
