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
