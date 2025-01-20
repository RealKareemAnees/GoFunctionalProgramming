# **Summary of all that shit**

## [Chapter 2]() _Functions as firts class citizens_

this chapter says how to treat functions as first class objects in go

### Type aliases

```go
// type aliases for primitive

type StringAlias string

func (s StringAlias) PrintS() {
	println(s)
}
```

- Now the `PrintS()` can operate on th s string, like in Javascript `string.topUppercase()` for example

```go
//usage

var s chapter2.StringAlias = "Hello, World!"
s.PrintS()
```

```go
// type aliases for functios

type Printer func(StringAlias)

func (p Printer) Print(s StringAlias){ // we also used the StringAlias
	p(s)// implamenting the functions
    // we can return the function again and do further chaining
}
```

```go
// usage
var p chapter2.Printer = func(s chapter2.StringAlias) {
		fmt.Println(s)
	}

p.Print("Hello, World!")
```

- we can also operate on _functions_

- **imagine anything else can be attached to types or structs**

### using functions as objects

- the core of first citizen

```go
// using functions as objects

type Adder func(int, int) int

func (a Adder) Add(x, y int) int {
	return a(x, y)
}

```

```go
// usage

var a chapter2.Adder = func(x, y int) int {
	return x + y
}

fmt.Println(a.Add(1, 2))
```

### next

everything else in this chapter is very basic programming and doesn't really worth mentioning

---

## [Chapter 3]() _Higher-Order Functions_

- **it is a function that either takes a function as input or return a function**

### partial application

as its name suggests, we execute the command on multiple stages, check code bellow

```go
// partial application

func Add(a, b, c int)  func(int) int {

    const r1 = a + b + c // first stage of exxecutions, it is successful and accually creates an output that satesifies the stage, but yet it is not usable

	return  func(d int) int {
		return r1 + d // this is the usable part
	}

}
```

```go
// usage of partial application

	add := chapter3.Add(1, 2, 3) // r1 = 6, we will be using this value now
	continueAdding := add(4)     // 10, it add 4 to the previous 6
	continueAdding2 := add(5)    // 11, it add 5 to the previous 6
	continueAdding3 := add(6)    // 12, it add 6 to the previous 6

	// print the results
	println(continueAdding)  // 10
	println(continueAdding2) // 11
	println(continueAdding3) // 12

```

### function currying

the idea of `function currying` is to chain the function but in higher order manner instead of passing a OOP-style struct (we will see that later on)

```go
func Combine(a, b, c int) int {
	return a + b + c
} // just normal function that adds 3 integers
```

below is the exact same function but using function currying

```go
func CombineWithCurrying(a int) func(int) func(int) int {
	return func(b int) func(int) int {
		return func(c int) int {
			return a + b + c
		}
	}
}
```

if you understand clousures then this thing is not black magic, look at the usage

```go
	// usage of currying
	combine := chapter3.CombineWithCurrying(1)(2)(3)  // 6
	combine2 := chapter3.CombineWithCurrying(1)(2)(4) // 7

	// print the results
	println(combine)  // 6
	println(combine2) // 7

```

**if you think about it, it is the same idea as _partial implementation_ but in a nested manner**

### next, i have skipped closures beacuese it doesn't really make any sense to explain it in functional programming summary
