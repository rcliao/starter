package main

import (
	"log"
	"net/http"

	"github.com/rcliao/starter/web"
)

func main() {
	client := web.NewClient()
	router := web.NewRouter(client)
	log.Fatal(http.ListenAndServe(":9000", router))
}
