package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/12sub/go-todo-app"
)

const (
	todoFile = ".todo.json"
)

func main() {
	// adding a command argument flag: just like argparse in python
	add := flag.Bool("add", false, "add a note in todo")
	complete := flag.Int("completed", 0, "completed todo list")
	del := flag.Int("del", 0, "delete a todo")
	list := flag.Bool("list", false, "showing todo list")

	// Parsing the flags
	flag.Parse()

	// getting the address / location of the Todos Struct from todo.go
	todos := &todo.Todos{}

	// Error handling: if the error equals the loaded todo struct &
	// the error value dosen't return nil, print the standard OS
	// error with the error value.
	// FPrintln formats using the default formats for its operands and writes to w
	if err := todos.Load(todoFile); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	// include switch statements for the argument flags
	switch {
	// argument flag "add": writing / adding a note
	// to the command line interface
	// store function stores every changes written to the
	// "todofile"
	case *add:
		task, err := getInput(os.Stdin, flag.Args()...)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}
		todos.Add(task)
		err = todos.Store(todoFile)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}
		// argument flag "complete": mark note status as
		// completed to the command line interface
		// store function stores every changes written to the
		// "todofile"
	case *complete > 0:
		todos.Completed(*complete)
		err := todos.Store(todoFile)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}
	case *del > 0:
		todos.Delete(*del)
		err := todos.Store(todoFile)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}
	case *list:
		todos.Print()
		// this is the default statements if every condition is not fufilled
	default:
		fmt.Fprintln(os.Stdout, "invalid command")
		os.Exit(0)
	}
}

func getInput(r io.Reader, args ...string) (string, error) {

	if len(args) > 0 {
		return strings.Join(args, " "), nil
	}

	scanner := bufio.NewScanner(r)
	scanner.Scan()
	if err := scanner.Err(); err != nil {
		return "", err
	}

	text := scanner.Text()
	if len(text) == 0 {
		return "", errors.New("empty todo isn't allowed")
	}

	return text, nil
}
