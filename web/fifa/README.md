# FIFA World Cup Winners
- This project exposes a Web API for accessing historic data from the FIFA World Cup championship.

## Running the server
- `go run server.go`

## Running tests
- `go test -v ./handlers/`

## Running with Docker
- To build the image from the Dockerfile
  - `docker build -t fifa .`
- To start an interactive shell, run:
  - `docker run -it --rm --name run-fifa fifa`
- From inside the shell, run the tests with:
  - `go test handlers/*`