package main

import (
	"fmt"
	"net/http"
)

type APIHandler struct{}

func (APIHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "kathir batch 11")
}
func main() {
	http.ListenAndServe(":8000", &APIHandler{})

}
