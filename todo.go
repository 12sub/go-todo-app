package cmd

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"time"
)

type item struct {
	Task        string
	Done        bool
	CreatedAt   time.Time
	CompletedAt time.Time
}

type Todos []item // converting the struct to an array

func (t *Todos) Add(task string) {
	// this is a function for adding tasks
	todo := item{
		Task:        task,
		Done:        false,
		CreatedAt:   time.Now(),
		CompletedAt: time.Time{},
	}

	*t = append(*t, todo)

}

func (t *Todos) Completed(index int) error {
	// this is a function for completed tasks
	list := *t

	if index <= 0 || index > len(list) {
		return errors.New("invalid index")
	}

	list[index-1].CompletedAt = time.Now()
	list[index-1].Done = true

	return nil
}

func (t *Todos) Delete(index int) error {
	// this is a function to delete tasks
	list := *t

	if index <= 0 || index > len(list) {
		return errors.New("invalid index")
	}

	*t = append(list[:index-1], list[index:]...)

	return nil
}

func (t *Todos) Load(filename string) error {
	// this is a function to load/ displays created tasks into the todo list
	file, err := ioutil.ReadFile(filename) // the calls a function that reads the file created
	if err != nil {                        // if there are still errors or if errors still exists
		if errors.Is(err, os.ErrNotExist) { //if the error tht exists cannot be verified
			return nil // return nothing
		} // else
		return err // return the error message
	}
	if len(file) == 0 { // if the length of the file is zero or there is no input in the created file
		return err // return an error message
	}

	err = json.Unmarshal(file, t) // Unmarshal parses encoded data and stores the result in the value pointed to by t
	if err != nil {               // typical error handling
		return err
	}

	return nil
}

func (t *Todos) Store(filename string) error {
	// This is a function to storing and writing tasks into the todo list
	data, err := json.Marshal(t) // this basically returns the JSON encoding of t
	if err != nil {
		return err
	}

	return ioutil.WriteFile(filename, data, 0644) // returns the written content
}
