package handler

import (
	"fmt"
	"net/http"
)

func Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, SayHello())
}

func SayHello() string {
	return "Hello World!"
}
