package main

import "fmt"

const englishHelloPrefix = "Hello, "

// Hello is the function to print hello

func Hello(name string) string {
	if name == "" {
		name = "World"
	}
	return englishHelloPrefix + name
}

func main() {
	fmt.Println(Hello("world"))
}
