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