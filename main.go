package main

import (
	"fmt"
	"fp/chapter2"
)

func main() {
	fmt.Println("Started main.go")

	// type aliases for primitive
	var s chapter2.StringAlias = "Hello, World!"
	s.PrintS()

	// type aliases for functions
	var p chapter2.Printer = func(s chapter2.StringAlias) {
		fmt.Println(s)
	}
	p.Print("Hello, World!")()

	// using functions as objects
	var a chapter2.Adder = func(x, y int) int {
		return x + y
	}
	fmt.Println(a.Add(1, 2))
}
