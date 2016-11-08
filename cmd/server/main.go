package main

import (
	"log"
	"net/http"

	"github.com/rcliao/starter/web"
)

func main() {
	client := web.NewClient()
	if err := client.Connect(); err != nil {
		log.Fatal(err)
	}
	router := web.NewRouter(client)
	log.Fatal(http.ListenAndServe(":9000", router))
}
