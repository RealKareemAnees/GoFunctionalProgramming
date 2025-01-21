package main

import "fmt"

func main() {

	fmt.Println(Fact(5))

}

func Fact(input int) int {
	if input == 0 {
		return 1
	}
	return (input) * Fact(input-1)
}
