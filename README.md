# Go
Golang for everything

- **Download and install** [here](https://golang.org/doc/install)
    - Verify installing: `go version`
    - Add the PATH: 
        - `sudo code ~/.zshrc` (if have VSCode)
        - `export PATH=$PATH:/usr/local/go/bin`
- **Setup `Go` workspace**
    - `mkdir $HOME/go && cd go`
    - `mkdir bin pkg src`
- If you store your code in a directory other than the default, you'll need to configure the `go tool` to look in the right place.
- Setup the `GOPATH`
    - `export GOPATH="/code"`
- Print Go environment information
    - `go env`
    - `go env [var ...]`
    - `go env -w [var ...]` --> set environment variable
## Packages
- The package names should be all lowercase, and ideally consist of single word.
- A package's functions can only be called from outside that package if they're exported. A function is exported if its name begins with a capital letter.
# Postgres
- Go inside the container
    -  `docker exec -it db bash`
- Go to localhose
    - `psql -h localhost -p 5432 -U postgres`