package main

import (
	"fmt"
	"net/http"
)

func Add(num1 int, num2 int) int {
	return num1 + num2
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func main() {
	fmt.Println("Apps started....")
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8181", nil)
}
