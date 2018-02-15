package main

import "fmt"

func Add(num1 int, num2 int) int {
	return num1 + num2
}

func main() {
	result := Add(1, 2)
	fmt.Println(result)
}
