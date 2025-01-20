package chapter3

// higher order functions

// partial application

func Add(a, b, c int) func(int) int {

	return func(d int) int {
		return a + b + c + d
	}

}

// currying

func Combine(a, b, c int) int {
	return a + b + c
}

func CombineWithCurrying(a int) func(int) func(int) int {
	return func(b int) func(int) int {
		return func(c int) int {
			return a + b + c
		}
	}
}
