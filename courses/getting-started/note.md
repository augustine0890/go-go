# Getting Started with Go
by [Coursera](https://www.coursera.org/learn/golang-getting-started/) - Univ of California, Irvine

---
## Week 1
- Complier: generates executable machine code for source code in a high-level language.
- Scope of a variable: region of a program where the variable can be accessed.
- Garbage Collection: deallocation of objects which are no longer in use.
- Interpreter: converts instructions in a high-level language into machine code at runtime
- Concurrency enables parallelism, but if parallel hardware is not present, speedup will not be achieved because multiple tasks cannot execute at the same time.
- In declaring a variable's type, the programmer specifies the amount of memory that needs to be allocated for that variable.
- The variable's type defines the nature of the variable along with the types of operations that can be performed on it.
- The variable's type specifies the nature of the data that can be contained (for example, integer vs. floating point number).

### Why Should I Learn Go?
**Advantages of Go**
- Code runs fast
- Garbage collection
- Simpler objects
- Concurrency is efficient

**Software Translation**
- Machine language: CPU instructions represented in binary
- Assembly language: CPU instructions with mnemonics
    - Easier to read
    - Equivalent to machine language
- High-level language: commonly used languages (C, C++, Java, Python, Go, etc.)
    - Much easier to use

- All software must be translated into the machine language of processor

**Compiled vs Interpreted**
- Complilation: translate instructions once before running the code
    - C, C++, Java (partially)
    - Translation occurs only once, saves time
- Interpretation: translate instructions while code is executed
    - Python, JS
    - Translation occurs every execution
    - Requires an interpreter

**Efficiency vs. Ease-of-Use**
- Compliled code is fast
- Interpreters make coding easier
    - Manage memory automatically
    - Infer variable types
- __Go__ is a good compromise

**Garbage Collection**
- Automatic memory management
    - Where should memory be allocated?
    - When can memory be deallocated?
- Manual memory management is hard
    - Deallocate too early, false memory accesses
    - Deallocate too late, wasted memory
- Go includes garbage collection
    - Typically only done bu interpreters

### Object-Orientd Programming
- Organize your code through encapsulation
- Group together data and functions are related
- User-defined type which is specific to an application
    - Ex. ints have data
- _Class_ defines data and fuctions
_ _Objects_ are instances of class

### Concurrency
**Performance Limits**
- Moore's Law used to help performance
    - Number of transistors doubles every 18 months
- More transistors used to lead to higher clock frequencies
- Power/temperature constraints limit clock frequencies now

**Parallelism**
- Number of cores still increases over time
- Multiple tasks may be performed at the same time on different cores
- Difficulties with parallelism
    - When do tasks start/stop?
    - What if one task needs data from another task?
    - Do tasks conflict in memory?

**Concurrent Programming**
- Concurrent is the management of multiple tasks at the same time
- Key requirement for large systems
- Concurrent programming enables parallelism
    - Management of task execution
    - Communication between tasks
    - Synchronization between tasks
- __Go__ includes concurrency primitives
- _Goroutines_ represent concurrent tasks
- _Channels_ are used to communicate between tasks
- _Select_ enables task synchronization

### Packages
- Group of related source files
- Each package can be imported by other packages
 
**Package Main**
- There must be one package called _main_
- Building the main package generated an executable program
- Main package needs a `main()` function
- `main()` is where code execution starts

### The GO Tool
- `go build`: compiles the program
    - Arguments can be a list of packages or a list of `.go` files
    - Creates an executable for the main package, same name as the first `.go` file.