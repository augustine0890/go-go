# Learn Go
> Eliminate the slowness and clumsiness of software development at Google, and thereby to make the process more productive and scalabel.

## Compiling
- Translate Go code to binary <- _compiler_ : a piece of software that converts Go code into a program that the computer understands.
- The transalted code is called an executable or a binary file.
- `go build` followed by the name of file -> `./file`
- The `go run` command combines both the compilation and execution of code.
- If the code is being used in production (users are interacting with it) then using an executable file if preferred.

## Basic Go Structure
**Packages**
- Programs that have the package declaration `package main` will create an executable file — it is a collection of code that can be used in other programs.
 - The _import statement_, `import "fmt"`. The import keyword allows us to bring in and use code from other packages.
- Packages serve important roles in Go. They group related code together, allow code to be reusable, and make it easier to maintain.
- A library is a Go package that won't create an executable file but contains code that we can use for our program.

**main**
- Functions: reusable blocks of code
    - The `func` keyword denotes the start of a function declaration.
    - `func` is followed by the name of the function.
    - After name is a pair of parentheses `()` and a set of curly braces `{}`.
- A `main` function is special, a file that has a `package main` declaration will automatically run the `main` function!

## Importing Multiple Packages
- [Here's a list of Go's standard packages](https://golang.org/pkg/)
- Single `import` with a pair of parentheses that contain our packages

## Comment
- Comments are ignored by the compiler and helpful for many things.
- A line comment is created using `//`
- The block comments that can span multiple lines— it starts with `/*`, ends with a `*/` and envelopes everything inside (including new lines):

## Go Resources
- Go also has it’s own built-in documentation system. 
- Use the command `go doc` followed by a package name. For instance, to find out more information on the `fmt` package: `go doc fmt`
- Get more specific information about a function in a package (like `fmt`‘s `Println` function) append `.Println` (or `.println`, the capitalization of the function doesn’t matter to `go doc`): `go doc fmt.Println`

## Values and Variables
- "Data" can be store by using variables

**Data Types**
- Integers, or `int`s, are whole/counting numbers. You would use an `int` to count the number of books on a shelf, the number of products in a warehouse, the number of people on a website, etc…
    - Basis Number Types in [Go](https://golang.org/ref/spec#Numeric_types)
- Floating-point numbers, or `float`s, can include fractional data. You would use a `float` to store distances, percentages, and other quantities that required division or precision.
- Complex numbers, `complex`, are pairs of floating-point numbers where the second part of the pair is marked with the “imaginary” unit `i`. Complex numbers are particularly useful when reasoning in 2-dimensional space and have other utilizations that make them relevant for involved calculations.
- Types also indicate how many bits (binary digits) will be used to represent the data. Fewer bits also means less data to save, so it will use less of a computer’s memory to hold onto that data.
- Signed integers can be negative, but unsigned integers can only be positive.
- `:=` create a new variable with the value appearing afterwards and infers type based on that value.

**fmt package**
- The verbs: `%v`, `%T`, `%d`, and `%f`.
- `fmt.Sprint()`, `fmt.Sprintln()`, `fmt.Sprintf()` will not print strings, but formats them.
- `fmt.Scan()` allows us to take in user input.

## Conditionals
- `if` statement that checks a condition and executes code if the condition is `true`
- `switch` statements can be used to check between multiple conditions much like an `if… else if… else` statement.
- Short variable declarations can be used prior to providing a condition for either `if` or `switch` statements. Declared variables are scoped to the statement blocks.
- The `math/rand` library’s `.Intn()` method is used to generate random numbers.
- Unique seed values can be created using time, namely `rand.seed(time.Now().UnixNano())`

## Functions
- When a similar pattern of code is used multiple times but with numbers of data tweaked slightly.
- `defer` call a function after the current function finishes

## Addresses and Pointers
- Go is a pass-by-value language
- The computer sets aside some space in its memory to store the value. The space that the computer allocates is called an _address_. Each address is marked as a unique numerical value.
- To find an address of a variable, use the `&` operator before a variable.
- Pointers are variables that specifically store addresses.
- A pointer is specific to what type of address it can store.
- The `*` operator can be used to assign a pointer the type of the value its address holds.
- The `*` operator can also be used to dereference a pointer and assign a new value.

## Select
- The `select` statement is used to choose from multiple send/receive channel operation.
- The `select` statement blocks until one of the send/receive operations is ready. If multiple operations are ready, one of them is chosen at random.