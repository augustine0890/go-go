# Sharing Work: Goroutines
- `Goroutines` let your program work on several different tasks at once.
- The goroutines can coordinate their work using `channels`, which let them send data to each other and synchronize so that one goroutine doesn't get ahead of another.
- Concurrency allows a program to pause one task and work on other tasks.
- Parallelism: running tasks simultaneously.
- A computer with only one processor can only run one task at a time.
- Concurrent tasks are called `goroutines`.
- `Goroutines` allow for concurrency: pausing one task to work on others. And in some situations they allow parallelism: working on multiple tasks simultaneously.

# Channels
- There is a way to communicate between goroutines --> `channels`.
- They ensure the sending goroutine has sent the value before the receiving goroutine attempts to use it.
