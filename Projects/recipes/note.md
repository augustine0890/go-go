# Application Project
```
.
├── go.mod
├── go.sum
├── handlers
│   └── handler.go
├── main.go
├── memory
│   └── main.go
├── models
│   └── recipe.go
├── note.md
├── recipes.json
└── swagger.json
```

# Go Swagger
- Generate the specification in JSON format
  - `swagger generate spec –o ./swagger.json`
- Load the generated spec in the Swagger UI
  - `swagger serve ./swagger.json`
- Set the flavor flag to `swagger`
  - `swagger serve -F swagger ./swagger.json`

# MongoDB Container
- Deploy a container
  - `docker run -d --name mongodb -e MONGO_INITDB_ROOT_USERNAME=admin -e MONGO_INITDB_ROOT_PASSWORD=password -p 27017:27017 mongo:4.4.3`
  - `-d` flag: detached mode
- Mount the directory to the `/data/db` directory
  - `docker run -d --name mongodb -v /Users/supertree/data:/data/db -e MONGO_INITDB_ROOT_USERNAME=admin -e MONGO_INITDB_ROOT_PASSWORD=password -p 27017:27017 mongo:4.4.3`
- Docker Secrets: 
  - `openssl rand -base64 12 | docker secret create mongodb_password`
  - Generate a random password for a MongoDB user and set it as a Docker secret.
  - `-e MONGO_INITDB_ROOT_PASSWORD_FILE=/run/secrets/mongodb_password`
- Remove container
  - `docker run -f container_name || true`
- Check the logs
  - `docker logs -f CONTAINER_ID`
- Stop and remove docker container
  - `docker stop container_id`
  - `docker container rm container_id`
  - `docker container prune`

# MongoDB Driver
- MongoDB Go Driver
  - `go get go.mongodb.org/mongo-driver/mongo`
- Pass `MONGO_URI`
  - `MONGO_URI="mongodb://admin:password@localhost:27017/test?authSource=admin" go run main.go`
- Pass `MONGO_DATABASE`
  - `MONGO_URI="mongodb://admin:password@localhost:27017/test?authSource=admin" MONGO_DATABASE=demo go run main.go`
  - `MONGO_URI="mongodb://admin:password@localhost:27017/test?authSource=admin" MONGO_DATABASE=demo go run *.go`
- Load `json` file directory into the collection
  - `mongoimport --username admin --password password --authenticationDatabase admin --db demo --collection recipes --file recipes.json --jsonArray`
