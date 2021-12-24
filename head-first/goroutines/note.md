# Sharing Work: Goroutines
- `Goroutines` let your program work on several different tasks at once.
- The goroutines can coordinate their work using `channels`, which let them send data to each other and synchronize so that one goroutine doesn't get ahead of another.
- Concurrency allows a program to pause one task and work on other tasks.
- Parallelism: running tasks simultaneously.
- A computer with only one processor can only run one task at a time.
- Concurrent tasks are called `goroutines`.
- `Goroutines` allow for concurrency: pausing one task to work on others. And in some situations they allow parallelism: working on multiple tasks simultaneously.
- New goroutines are started with a `go` statement: an ordinary function call preceded by the `go` keyword.

# Channels
- There is a way to communicate between goroutines --> `channels`.
- They ensure the sending goroutine has sent the value before the receiving goroutine attempts to use it.
- By default, sending a value on a channel blocks (pauses) the current goroutine until that value is received. Attempting to receive a value also blocks the current goroutine until a value is sent on that channel.
- Channels are created by calling the built-in `make` function.
- Each channel only carries values of one particular type.
- Channels can carray composite types like slices, maps, and strucs.