# Concurrency

## Concurrency and parallelism
- Single task: have a single task
- Concurrency: have multiple tasks (Don't need at the same time)
- Parallelism: execute multiple tasks simultaneously
- Concurrency is a prerequisite for parallel execution

## Goroutines
- Thread:
  - Ability of OS to run the tasks
  - Have own execution stack
  - Fixed stack space (around 1MB)
  - Managed by OS

- Goroutines:
  - Have own execution stack
  - Variable stack space (starts @2KB)
  - Manage by Go runtime

## The `Sync` Package
- Coordinating tasks: `WaitGroups`

`sync.WaitGroup`
- A WaitGroup waits for a collection of goroutines to finish.

- Shared memory: `Mutexes`
- `Mutex`: a mutal exclusion lock, ensurses that only one task can be accessing the mutex at one time.
- `RWMutex`: allows multiple readers to access the protected code, but only one task can write to mutex.