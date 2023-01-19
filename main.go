package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", HelloHandler)
	http.ListenAndServe(":8080", nil)
}

func HelloHandler(w http.ResponseWriter, _ *http.Request) {
	foo()
	fmt.Fprintf(w, "Hello World!")
}

func foo() {
	fmt.Println("foo")
}
