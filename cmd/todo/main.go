package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/12sub/go-todo-app"
)

const (
	todoFile = ".todo.json"
)

func main() {
	add := flag.Bool("add", false, "add a note in todo")
	flag.Parse()
	todos := &todo.Todos{}

	if err := todos.Load(todoFile); err != nil {
		fmt.Println(os.Stderr, err.Error())
		os.Exit(1)
	}
	switch {
	case *add:
		todos.Add("Subomi's Todo")
		err := todos.Store(todoFile)
		if err != nil {
			fmt.Println(os.Stderr, err.Error())
			os.Exit(1)
		}
	default:
		fmt.Fprintln(os.Stdout, "invalid command")
		os.Exit(0)
	}
}
