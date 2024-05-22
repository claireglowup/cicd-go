package main

import (
	"ci-cd/handler"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler.Hello)

	log.Println("server start...")

	http.ListenAndServe(":8080", nil)
}
