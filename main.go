package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

var stdOutLog = log.New(os.Stdout, "Access: ", log.Llongfile)
var stdErrLog = log.New(os.Stderr, "Error: ", log.Llongfile)

func Add(num1 int, num2 int) int {
	return num1 + num2
}

type Person struct {
	Name string
	Age  int
}

func handler(w http.ResponseWriter, r *http.Request) {
	stdOutLog.Println("Ini access log")
	stdErrLog.Println("Ini error log")
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
	handler(w, r)
}

func main() {
	fmt.Println("Apps started....")
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8181", nil)
}
