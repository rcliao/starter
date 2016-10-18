package web

import (
	"fmt"
	"net/http"
)

// Index renders the hello world message
func Index(client *Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Print(w, "Hello GoLang web starter!")
	}
}
