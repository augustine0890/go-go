# Concurrencty Patterns
## Runner
- The purpose of the runner package is to show channels can be used to monitor the amount of time a program is running and terminate the program if it runs too long.
- You can use channels to control the lifetime of programs.
- A `select` statement with a default case can be used to attempt a nonblocking send or receive on channel.