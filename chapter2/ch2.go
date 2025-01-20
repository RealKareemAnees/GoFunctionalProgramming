package chapter2

// type aliases for primitive

type StringAlias string

func (s StringAlias) PrintS() {
	println(s)
}

// type aliases for functios

type Printer func(StringAlias)

func (p Printer) Print(s StringAlias) func() {
	p(s)
	return s.PrintS
}

// using functions as objects

type Adder func(int, int) int

func (a Adder) Add(x, y int) int {
	return a(x, y)
}
