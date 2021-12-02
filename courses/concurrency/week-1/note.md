# Why Use Concurrency
## Parallel Execution
- Two programs execute the parallel if they execute at exactly the same time.
- Need replicated hardware.
- Some tasks are parallelizable and some are not.

## Von Neumann Bottleneck
- Get speedup without changing software by design faster processors.
- CPU <-> Memory
- Design processor with more memory
  - Cache access time = 1 clock cycle
  - Main memory access time = ~100 clock cycle
  - Increasing on-chip cache improves performance
- Moore's Law
  - Smaller transistors switch faster
  - Exponential increase in density would lead to exponential increase in speed.

## Power Wall
- Increasing transistors density leads to increased power consumption
- High power leads to high temperature

## Multi-Core Systems
- Parallel execution is needed to exploit multi-core systems
- Code made to execute on multiple cores
- Different programs on different cores

## Concurrent Execution
- Concurrent execution is not necessarily the same as parallel execution.
- Concurrent: start and end times overlap
- Parallel: execute at exactly the same time
- Parallel tasks must be executed on different hardware
- Concurrent tasks may be executed on same hardware
- Concurrency improves performance, even without parallelism
- Tasks must periodically wait for something
- Other concurrent tasks can operate while one task is waiting
