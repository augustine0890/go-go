# `weight-tracker`

- Create the module
    - `go mod init weight-tracker`

## Application Structure
```
weight-tracker
- cmd
  - server
    main.go
- pkg
  - api
    user.go
    weight.go
  - app
    server.go
    handlers.go
    routes.go
  - repository
    storage.go
```
- The `cmd` is the entry point into the program and gives the flexibility to interact with the program.
- The `pkg` contain everything else: routes, database interactions, services
- There are four packages: main, api, app, repository