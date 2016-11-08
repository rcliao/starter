package web

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/rcliao/starter"
)

// Index renders the hello world message
func Index(client *Client) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello GoLang web starter!")
	})
}

// GetTodos returns all todos in database
func GetTodos(client *Client) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		list, err := client.todoService.List()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintln(w, "ERROR")
			return
		}
		log.Printf("%v", list)
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(list)
	})
}

// AddTodo is a post request to add a new todo
func AddTodo(client *Client) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		decoder := json.NewDecoder(r.Body)
		var t starter.Todo
		err := decoder.Decode(&t)
		if err != nil {
			panic("Oh no, right? OH NO PANIC!")
		}
		client.todoService.Add(t.Name, t.Description)
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, "Done")
	})
}
