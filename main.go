package main

import (
	"fp/chapter3"
)

func main() {

	// usage of partial application
	add := chapter3.Add(1, 2, 3) // 6
	continueAdding := add(4)     // 10, it add 4 to the previous 6
	continueAdding2 := add(5)    // 11, it add 5 to the previous 6
	continueAdding3 := add(6)    // 12, it add 6 to the previous 6

	// print the results
	println(continueAdding)  // 10
	println(continueAdding2) // 11
	println(continueAdding3) // 12

	// usage of currying
	combine := chapter3.CombineWithCurrying(1)(2)(3)  // 6
	combine2 := chapter3.CombineWithCurrying(1)(2)(4) // 7

	// print the results
	println(combine)  // 6
	println(combine2) // 7

}
