# **Summary of the book [Functional Programming in Go](https://www.oreilly.com/library/view/functional-programming-in/9781801811163/)**

![Book Cover](image.png)

<i>
<h3 style="color: white; font-family: 'Times New Roman', serif;">
I have read this book and I highly recommend it to anyone in the software engineering industry. I personally think developers should learn FP before OOP. This summary is meant for people who want a refresher or those who already know FP and want to see it applied in Go.
</h3>
<h3 style="color: white; font-family: 'Times New Roman', serif;">
If this is your first time with FP, please read this awesome 10/10 book from <a href="https://www.oreilly.com/library/view/functional-programming-in/9781801811163/">O'Reilly</a>.
</h3>
</i>

> I have written this summary to cover everything as much as possible. If anything is skipped, I have mentioned, "I have skipped X topic for reason Y."

> If you like this summary and are interested in more, you can follow me on [LinkedIn](https://www.linkedin.com/in/kareem-anees-0496b62b3) or [Facebook](https://www.facebook.com/TheyCallMeAbdo).

---

<br>

---

# Table of Contents

- [**Summary of the Book: Functional Programming in Go**](https://www.oreilly.com/library/view/functional-programming-in/9781801811163/)
  - [Content Table](#content-table)
- [**Part 1: Essentials**](#part-1-essentials)
  - [Chapter 2: Functions as First-Class Citizens](#chapter-2-functions-as-first-class-citizens)
    - [This chapter explains how to treat functions as first-class objects in Go](#this-chapter-explains-how-to-treat-functions-as-first-class-objects-in-go)
    - [Type Aliases](#type-aliases)
    - [Using Functions as Objects](#using-functions-as-objects)
    - [Other Basics](#other-basics)
  - [Chapter 3: Higher-Order Functions](#chapter-3-higher-order-functions)
    - [What is a Higher-Order Function?](#what-is-a-higher-order-function)
    - [Partial Application](#partial-application)
    - [Function Currying](#function-currying)
    - [Closures](#closures)
  - [Chapter 4: Purity and Testability](#chapter-4-purity-and-testability)
    - [What is Purity?](#what-is-purity)
    - [Referential Transparency](#referential-transparency)
    - [Use Cases for Pure and Impure Functions](#use-cases-for-pure-and-impure-functions)
  - [Chapter 5: Immutability](#chapter-5-immutability)
    - [Garbage Collection, Escape Analysis, Stack, and Heap](#garbage-collection-escape-analysis-stack-and-heap)
    - [Functor](#functor)
    - [Monads](#monads)
- [**Part 2: Using FP Techniques**](#part-2-using-fp-techniques)
  - [Chapter 6: Common Categories of Functions](#chapter-6-common-categories-of-functions)
    - [Predicate-Based Functions](#predicate-based-functions)
    - [TakeWhile & DropWhile vs. Filter](#takewhile-dropwhile-vs-filter)
    - [Transformers/Maps](#transformersmaps)
    - [Reducers](#reducers)
  - [Chapter 7: Recursion](#chapter-7-recursion)
    - [Recursion vs. Loops](#recursion-vs-loops)
  - [Chapter 8: Composition of Functions](#chapter-8-composition-of-functions)
    - [Dot Notation Chaining](#dot-notation-chaining)
    - [Continuation-Passing Style (CPS)](#continuation-passing-style-cps)
- [**The End**](#the-end)

---

<br>

# **PART 1: _ESSENTIALS_**

Basic concepts that require some attention.

## [Chapter 2]() _Functions as First-Class Citizens_

This chapter explains how to treat functions as first-class objects in Go.

---

### Type Aliases

```go
// Type aliases for primitive

type StringAlias string

func (s StringAlias) PrintS() {
    println(s)
}
```

<br>

- Now the `PrintS()` can operate on the `s` string, similar to how `string.toUpperCase()` works in JavaScript.

```go
// Usage

var s chapter2.StringAlias = "Hello, World!"
s.PrintS()
```

```go
// Type aliases for functions

type Printer func(StringAlias)

func (p Printer) Print(s StringAlias) { // We also used the StringAlias
    p(s) // Implementing the function
    // We can return the function again and do further chaining
}
```

```go
// Usage
var p chapter2.Printer = func(s chapter2.StringAlias) {
    fmt.Println(s)
}

p.Print("Hello, World!")
```

- We can also operate on _functions_.

- **Imagine anything else can be attached to types or structs.**

---

### Using Functions as Objects

- The core of first-class citizens.

```go
// Using functions as objects

type Adder func(int, int) int

func (a Adder) Add(x, y int) int {
    return a(x, y)
}
```

```go
// Usage

var a chapter2.Adder = func(x, y int) int {
    return x + y
}

fmt.Println(a.Add(1, 2))
```

---

Everything else in this chapter is basic programming and doesn’t really warrant mentioning.

---

## [Chapter 3]() _Higher-Order Functions_

**A higher-order function is a function that either takes a function as input or returns a function.**

### Partial Application

As its name suggests, we execute the command in multiple stages. Check the code below.

```go
// Partial application

func Add(a, b, c int) func(int) int {
    const r1 = a + b + c // First stage of execution

    return func(d int) int {
        return r1 + d // This is the usable part
    }
}
```

```go
// Usage of partial application

add := chapter3.Add(1, 2, 3) // r1 = 6, we will use this value now
continueAdding := add(4)     // 10, adds 4 to the previous 6
continueAdding2 := add(5)    // 11, adds 5 to the previous 6
continueAdding3 := add(6)    // 12, adds 6 to the previous 6

// Print the results
println(continueAdding)  // 10
println(continueAdding2) // 11
println(continueAdding3) // 12
```

---

### Function Currying

The idea of `function currying` is to chain the function but in a higher-order manner instead of passing an OOP-style struct (we will see that later on).

```go
func Combine(a, b, c int) int {
    return a + b + c
} // Just a normal function that adds 3 integers
```

Below is the exact same function but using function currying.

```go
func CombineWithCurrying(a int) func(int) func(int) int {
    return func(b int) func(int) int {
        return func(c int) int {
            return a + b + c
        }
    }
}
```

If you understand closures, then this thing is not black magic. Look at the usage:

```go
// Usage of currying
combine := chapter3.CombineWithCurrying(1)(2)(3)  // 6
combine2 := chapter3.CombineWithCurrying(1)(2)(4) // 7

// Print the results
println(combine)  // 6
println(combine2) // 7
```

## **If you think about it, it is the same idea as _partial application_ but in a nested manner.**

I have skipped closures because it doesn’t really make any sense to explain it in a functional programming summary.

---

## [Chapter 4]() _Purity and Testability_

<h3 style="color: red;">Important!</h3>

I know you might skip this part, but it’s worth mentioning two things: first, there are some useful topics here; second, escape analysis.

> Pure functions are essential in Go in terms of performance and **bottlenecks**. If you want to learn more about it, visit my post about how purity can skyrocket your Go code performance [here](https://www.facebook.com/share/p/188hthkfx9/). Also, see this blog about how crazy GC can slow down your app [here](https://medium.com/@roopa.kushtagi/netflixs-journey-from-g1gc-to-zgc-tackling-tail-latencies-66ee75681d49).

---

### What is Purity?

_Purity is not making any side effects._ Also, purity is coupled with immutability.

- **Mutable**: A mutable data structure can be modified after it has been constructed.
- **Immutable**: Cannot be modified; it can only be deleted and recreated.
- **Pure Function**: A function that doesn’t change any state during the program's lifetime.

> **_The keyword `const` in Go creates an immutable data structure that cannot be altered after its construction. Also, the GC doesn’t check any `const` declared data structures in the heap._**

---

### Referential Transparency

It is when we can replace the function with its output.

**In other words, if we use the same function with the same inputs multiple times, it gives the same output (_statelessly_ without side effects or states).**

An example:

```go
func a(x, y int) int {
    return x + y
} // Referentially transparent as it always returns the same x + y for the same inputs

func b() int {
    return time.Now().Second()
} // Returns a unique output every time (not referentially transparent)
```

> **Idempotency**: It is the same concept as _referential transparency_; a function returns the same outputs with the same inputs every time it is called.

---

The rest of this chapter discusses the benefits and use cases of pure and impure functions, which are very clear and not worth mentioning.

---

## [Chapter 5]() _Immutability_

<p style="color:red;">The first part of this chapter talks about GC, escape analysis, stack, and heap. These are more of a prerequisite for software engineering in general rather than an FP concept.</p>

### Functor

A functor is a way to apply function `a` on data structure `x`.

```go
// Functors

func FunctorAX[A, B any](strategy func(A) B, x []A) []B {
    var result []B
    for _, v := range x {
        result = append(result, strategy(v))
    }
    return result
}
```

```go
// Usage

func strategy(a int) int { return a * 2 }

var x = []int{1, 2, 3, 4, 5}
var y = FunctorAX(strategy, x) // [2, 4, 6, 8, 10]
```

- The function `strategy` doesn’t know anything except that it takes `a` and returns `a * 2`.

- `FunctorAX` takes a set of elements (it doesn’t have to be of type `A`) and loops over them, calling `strategy` on each of them.

**Another simpler example:**

```go
func Functor2(strategy func(any) any, x any) any {
    return strategy(x) // Clear, huh?
}
```

---

### Monads

<p style="color:magenta;">
If you don’t understand the following explanation, watch this <a href="https://www.youtube.com/watch?v=HIBTu-y-Jwk&t=404s">video</a>. Monads are quite tricky to wrap your head around.
</p>

- I like to think of **monads** as _partial application_ but with two more concepts:

  1.  `Maybe`: Maybe the data structure we want is empty, so we wrap it in a `Maybe` data structure that has a value.

  2.  `Unit()`: Just a function that returns the `Maybe` instead of direct interaction with it.

  3.  `Bind()`: To chain operations on this monad, use `Bind`. `Bind` is a function that takes another function as input, applies it to the monad, and returns `Maybe` so further chaining can be done if needed.

Look at this code:

```go
// Create Maybe to wrap an integer value
type Maybe struct {
    value *int
}

// Unit wraps a value in a Maybe box.
func Unit(val int) Maybe {
    return Maybe{value: &val}
}

// Bind function that operates on Maybe
func (m Maybe) Bind(f func(int) Maybe) Maybe {
    if m.value == nil {
        return Maybe{value: nil}
    }

    return f(*m.value) // Apply the function
}

// Example function to double a number
func Double(n int) Maybe {
    return Unit(n * 2)
}

func main() {
    // Start with a number in a Maybe box
    result := Unit(5). // Returns a Maybe with value 5
        Bind(Double). // Double it -> 10
        Bind(Double). // Double it again -> 20
        Bind(Double)  // Double it again -> 40
}
```

---

<br>

<br>

# **PART 2: USING FP TECHNIQUES**

It is becoming more exciting! 🔥

> If you have reached this point, I want to really thank you for your attention and time. I spent almost 3 hours on **Part 1**.

---

## [Chapter 6]() _Common Categories of Functions_

### Predicate-Based Functions

A predicate-based function is a function that returns either `true` or `false`.

```go
type Predicate[T any] func(T) bool
```

```go
// A predicate function
func isGreaterThan0(val int) bool {
    return val > 0
}

// Common use case
func functor(val int, compare Predicate[int]) bool {
    // Perform work
    return compare(val)
}
```

---

### TakeWhile & DropWhile **VS** Filter

Let us say we want to check if a list has a value divisible by 2 using predicates. An example predicate would be:

```go
func isDivisibleBy2(i int) bool {
    return i%2 == 0
}
```

<br>

`takeWhile()` loops over elements, applies a condition, and once the condition fails, it returns the elements processed before the failure.

```go
// Expressive code from the book:
func TakeWhile[A any](input []A, pred Predicate[A]) []A {
    out := []A{}
    for _, element := range input {
        if !pred(element) {
            return out
        }
        out = append(out, element)
    }
    return out
}

// Usage
ints := []int{1, 1, 2, 3, 5, 8, 13}
result := TakeWhile(ints, isDivisibleBy2) // [1, 1], condition is met at 2
```

<br>

`dropWhile()`, on the other hand, skips elements while a condition holds true and returns the remaining elements once the condition fails.

```go
// Expressive code from the book:
func DropWhile[A any](input []A, pred Predicate[A]) []A {
    out := []A{}
    drop := true
    for _, element := range input {
        if pred(element) {
            drop = false
        }
        if !drop {
            out = append(out, element)
        }
    }
    return out
}

// Usage
ints := []int{1, 1, 2, 3, 5, 8, 13}
result := DropWhile(ints, isDivisibleBy2) // [2, 3, 5, 8, 13]
```

<br>

`Filter` applies a condition to every element in the list and returns a list of elements that satisfy the condition. More details on filters are provided later in this document.

> ## **Common Uses of `dropWhile` and `takeWhile`**
>
> 1. **Processing a Stream of Data**  
>    Skip or extract specific sections of sequential data efficiently without loading it all into memory.
>
> 2. **Finding a Subsection of a List**  
>    Extract dynamic segments of data based on increasing or decreasing values.
>
> 3. **Parsing Input Until a Condition is Met**  
>    Efficiently process structured or semi-structured text with delimiters.
>
> 4. **Filtering Events in Time Series Data**  
>    Manage ranges of interest in time-series data based on thresholds.
>
> 5. **Handling Sorted Data**  
>    Work with sorted sequences where conditions are monotonic.
>
> 6. **Batch Processing in Pipelines**  
>    Handle varying priorities or task states dynamically.
>
> 7. **Cleaning Data**  
>    Automate operations such as removing leading zeros or anomalies.

---

### Transformers / Maps

Transformers or `maps` (not `hashMaps`) are higher-order functions that apply a transformation to elements of a data structure.

```go
type MapFunc[A any] func(A) A

func Map[A any](input []A, m MapFunc[A]) []A {
    output := make([]A, len(input))
    for i, element := range input {
        output[i] = m(element)
    }
    return output
}
```

> The above `Map()` function is a _pure Map_, which doesn't modify the original slice but creates a new one. This is often preferable in functional programming.

---

### Reducers

Reducers combine all elements of a data structure into a single value using a function.

```go
// Reducer function type
type ReduceFunc[T any] func(p1, p2 T) T
```

```go
func Reduce[A any](input []A, reducer ReduceFunc[A]) A {
    result := input[0] // Start with the first element
    for _, element := range input[1:] { // Iterate from the second element
        result = reducer(result, element)
    }
    return result
}
```

```go
// Usage example
func sum(a1, a2 int) int { return a1 + a2 }

total := Reduce([]int{1, 2, 3, 4, 5}, sum) // 1 + 2 + 3 + 4 + 5 = 15
```

> _Reducers_ can significantly improve time complexity in certain scenarios.

---

## [Chapter 7]() _Recursion_

_Recursion_ occurs when a function calls itself. It is particularly useful for problems like tree traversal but can also increase complexity and risk stack overflows if not used carefully.

```go
func recursive() {
    if condition { // Base case
        return
    }
    recursive() // Recursive case
}
```

#### Example: Factorial Calculation

```go
func Fact(input int) int {
    if input == 0 {
        return 1
    }
    return input * Fact(input-1)
}

// Usage
Fact(5) // 120
```

The computation can be visualized as:

```
Fact(5) = 5 * Fact(4) = 5 * 24 = 120
Fact(4) = 4 * Fact(3) = 4 * 6 = 24
Fact(3) = 3 * Fact(2) = 3 * 2 = 6
Fact(2) = 2 * Fact(1) = 2 * 1 = 2
Fact(1) = 1 * Fact(0) = 1 * 1 = 1
Fact(0) = 1
```

---

> Recursion is not recommended in production unless absolutely necessary. Loops are generally more efficient and easier to debug in most cases.

---

## [Chapter 8]() _Composition of Functions_

This chapter focuses on chaining and continuous execution of functions in a readable and functional manner.

---

### Dot Notation Chaining

In **_Go_**, you can attach methods to types to enable chaining:

```go
package main

type Str struct {
    value int
}

func (s *Str) add(x int) *Str {
    s.value += x
    return s
}

func (s *Str) subtract(x int) *Str {
    s.value -= x
    return s
}

func main() {
    s := &Str{value: 10}
    s.add(5).subtract(3)
    println(s.value) // 12
}
```

For simpler cases, define a custom type directly:

```go
type number int

func (n number) add(x number) number {
    return n + x
}

func (n number) sub(x number) number {
    return n - x
}

func main() {
    var n number = 5
    n = n.add(3).sub(2)
    println(n) // 6
}
```

---

### CPS (Continuation-Passing Style)

CPS involves passing the "next step" as a callback function. It is widely used in event-driven architectures for non-blocking execution.

```go
import "sync"

var wg sync.WaitGroup

func APICall(url string, cb func(string)) {
    wg.Add(1)
    go func() {
        defer wg.Done()
        cb("Data from API call: " + url)
    }()
}

func main() {
    APICall("http://example.com", func(data string) {
        println(data + " API call 1")
    })
    APICall("http://example.com", func(data string) {
        println(data + " API call 2")
    })
    wg.Wait() // Wait for all API calls to finish
}
```

---

<br>

# The End

Thanks for reading!

> **But where is _Part 3_???**  
> Well, Part 3 is just a brief overview of common design patterns and libraries. This book excels in explaining the basics, but that part is another topic for another day. Also, I’m not yet confident enough to summarize _design patterns_!
