package main

type Todo struct {
	ID        int
	Title     string
	Completed bool
}

var nextTodoId = 3

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
