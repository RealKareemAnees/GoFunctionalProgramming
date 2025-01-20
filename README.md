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

### everything else in this chapter is very basic programming and doesn't really worth mentioning

---

## [Chapter 3]() _Higher-Order Functions_

**it is a function that either takes a function as input or return a function**

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

### i have skipped closures beacuese it doesn't really make any sense to explain it in functional programming summary

---

## [Chapter 4]() _purity and testability_

<h3 style= "color: red;" >Important!</h3>

i know that you are going to skip this part but it worth mentioning two things; first: their are some useful topics here, second: escape analysis

> pure functions are essensial in go in terms of performance and **bottlenecks**, if you want to learn more about it visit my post about how purity can highrocket yor golang code performance [here](https://www.facebook.com/share/p/188hthkfx9/), also see this blog about how crazy can GC slow down your app [here](https://medium.com/@roopa.kushtagi/netflixs-journey-from-g1gc-to-zgc-tackling-tail-latencies-66ee75681d49)

### what is purity

_purity is not making any side effect_, also purity is coupled with immutability

- **mutable**: a mutable DS can be modified after it has constructed
- **immutable**: cannot be modified, it can only be deleted and recreated
- **Pure function**: a function that doesn't change any state in the program lifetime

> **_the keyword `const` in go creates an immutable DS and it cannot be altered after it's construction, Also the GC doesn't check any `const` declared DS in the heap_**

### Refrential transparency

it is when we can replace the function with it's output

**in other words, if we use the same function with the same inputs multiple times, each time it gives the same output (_statelessly_ without side effects or states)**

an example

```go
func a(x,y int) {
	return x + y
} // Refrentially treansparent as it always return the same x + y for the same x + y

func b() {
	return time.now()
} // every moment it returns a unique output (not Refrentially treansparent)
```

> **Idempotency**: it is the same concept as _Refrential transparency_; a function returns the same outputs with the same inputs everytime it is called

### the rest of this chapter talks about the benifets and use cases of pure and impure functions which is very clear actually and is not worth mentioning
