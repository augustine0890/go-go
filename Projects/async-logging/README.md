__Build a fully asynchronous logging package__

How to run the program
- Navigate into the `./cmd` directory
- Run as the defaults
  - `go run .`
  - The application will run in asynchronously mode and write it's output directly to the shell
    - The `-out` flag allows a destination file to be specified
    - The `-async` flag determines if the logger will asynchronously or not. It is async by default, which will not work properly in the application's initial condition.

To run the test
- Use the `go test` command from the command line
- To focus on a specific test, use the `-run` flag followed by a pattern that matches the test name
  - e.g. `go test -run MessageChannel` will run the first test
  