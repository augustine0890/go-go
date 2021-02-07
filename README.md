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

# Postgres
- Go inside the container
    -  `docker exec -it db bash`
- Go to localhose
    - `psql -h localhost -p 5432 -U postgres`