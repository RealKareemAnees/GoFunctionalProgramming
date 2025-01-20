package chapter5

// functors

func functotAX[a, b any](strategy func(a) b, x []a) []b {
	var result []b
	for _, v := range x {
		result = append(result, strategy(v))
	}
	return result
}

//usage

func strategy(a int) int { return a * 2 }

var x = []int{1, 2, 3, 4, 5}
var y = functotAX(strategy, x) // [2, 4, 6, 8, 10]

func functor2(strategy func(any) any, x any) any {
	return strategy(x)
}

// monads
