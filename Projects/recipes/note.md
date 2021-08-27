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
  - `docker run -d --name mongodb -v /Users/supertree/data/mongo:/data/db -e MONGO_INITDB_ROOT_USERNAME=admin -e MONGO_INITDB_ROOT_PASSWORD=password -p 27017:27017 mongo:4.4.3`
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

# Redis Docker
- Run a container
  - `docker run -d --name redis -p 6379:6379 redis:letest`
  - The `-d` flag runs the Redis container as a daemon
  - The `-p` flag maps port 6379 of the container to port 6379 of the host
- `redis.conf` file
  - `maxmemory-policy allkeys-lru` --> Least Recently Used (LRU) deletes the cache items that were the least recently used.
  - `maxmemory 512mb`
- [Redis Insights](https://redislabs.com/fr/redis-enterprise/redis-insight/) (GUI)

# Performance Benchmark
- Simulate multiple request with [Apache Benchmark](https://httpd.apache.org/docs/2.4/programs/ab.html)
  - `ab -n 2000 -c 100 -g without-cache.data http://localhost:8080/recipes`
  - `ab -n 2000 -c 100 -g with-cache.data http://localhost:8080/cache/recipes`
- Time taken for tests: total time to complete the 2000 requests
- Time per request: how many milliseconds it takes to complete one request.

# Authentication
- Generate a random secret string
  - `openssl rand -base64 16`
- JWTs (JSON Web Tokens)
  - `go get github.com/dgrijalva/jwt-go`
- Sign-in using a username/password --> valid <---> JWT token. The client will use the JWT token in future request (`Authorization` header)

# Client sessions and cookies
- Session cookies allow users to be recognized within an application without having to authenticate. Without cookies, every time you issue an API request, the server will treat you like a completely new visitor.
- Gin middleware for session
  - `go get github.com/gin-contrib/sessions`

# HTTPS Server
- Download [Ngrok](https://ngrok.com/download) at `https://ngrok.com/download`
- Uzip: `unzip ngrok-stable-darwin-amd64.zip`
- `cp ngrok /usr/local/bin`
- `chmod +x /usr/local/bin/ngrok`
- Verify installation: `ngrok version`
- Configure Ngrok to listen and forward requests into port 8080
  - `ngrok http 8080`

# Self-signed certificates
- SSL certificates are what websites use to move from HTTP and HTTPs. The certificate uses SSL/Transport Layer Security (TLS) encryption to keep user data secure, verify ownership of the website.
- Two files generate
  - `localhost.crt`: self-signed certificate
  - `localhost.key`: private key
```
mkdir certs
openssl req -x509 -nodes -days 365 -newkey rsa:2048 -keyout certs/localhost.key -out certs/localhost.crt
```

# Single-Page Applications (SPA)
- Install `create react app`
  - `npm install -g creates-react-app`
- Create React project
  - `create-react-app <recipes-client>`