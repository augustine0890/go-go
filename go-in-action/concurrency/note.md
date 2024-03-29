# Concurrency
- Concurrency synchronization comes from a paradigm called "communication sequential processes" or "CSP".
- CSP is a message-passing model that works by cummunicating data between goroutines instead of locking data to synchronize access.

## Concurrency vs Parallelism
- `Process` like a container that holds all the recources an application uses and maintains as it runs.
- The resources include but are not limited to a memory address space, handles to files, devices, and threads.
- A `thread` is a path of execution that's scheduled by the operating system to run the code that you write in your functions.
- Each process contains at least one thread, and the initial thread for each process is called the _main thread_.
- The Go runtime schedules goroutines to run against logical processors. Each logical processor is individually bound to a single operating system thread. Even with a single logical processor, hundreds of thousands of goroutines can be scheduled to run concurrently with amazing efficiency and performance.
- Parallelism can only achieved when multiple pieces of code are executing simultaneously against different physical processors.
- Parallelism is about doing a lot of things at once. Concurrency is about managing a lot of things at once.
- If you want to run goroutines in parallel, you must use more than one logical processor. When there are multiple logical processors, the scheduler will evenly distributed goroutine between the logical processors --> goroutines will be running different threads.

## Race conditions
- When two or more goroutines have unsynchronized access to a shared resources and attempt to read and write to that resource at the same time --> `race condition`.
- Race detector: `go run -race main.go`

## Locking shared resources
- Atomic functions provide low-level locking mechanism for synchronizing access to integers and pointers.

## Mutexes
- A mutex is named after the concept of mutual exclusion.
- A mutex is used to create a critical section around code that ensures only one goroutine at a time can excute that code section.

## Channels
- Channels provide an intrinsic way to safely share data between two goroutines.
- An _unbuffered channel_ is a channel with no capacity to hold any value before it's received.
- A _buffered channel_ is a channel with a capacity to hold one or more values before they're received.
- An unbuffered channel provides a guarantee that an exchange between two goroutines is performed at the instant the send and receive take place. A buffered channel has no such guarantee.