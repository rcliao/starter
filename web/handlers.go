package web

import (
	"fmt"
	"net/http"
)

// Index renders the hello world message
func Index(client *Client) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello GoLang web starter!")
	})
}
