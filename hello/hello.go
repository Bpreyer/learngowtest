package main

import "fmt"

// Hello is the function to print hello

func Hello(name string) string {
	return "Hello, " + name
}

func main() {
	fmt.Println(Hello("world"))
}
