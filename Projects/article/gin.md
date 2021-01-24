# Gin
- Gin is a high-performance micro-framework that can be used to build web applications and microservices.

## Designing the Application
- The control flow for a typical web app, API server or a microservice:
```
Request --> Route Parser --> [Optional Middleware] --> Route Handler --> [Optional Middleware] --> Response
```

## Application Functionality
- Building a simple article manager:
    - Users register, login, log-out
    - Create new articles
    - List of all articles on the home page
    - Display a single article on its own page
- The functionalities offered by Gin:
    - Routing: to handle various URLs
    - Custom rendering: to handle the response format
    - Middleware: to implement authentication

## Routing
- Serve the index page at route - HTTP GET request
- Group user-related routes under the `/u/` route:
    - Serve the login page at `/u/login` (HTTP GET request)
    - Process the login credentials a `/u/login` (HTTP POST request)
    - Log out at `/u/logout` (HTTP GET request)
    - Serve the registration page at `/u/register` (HTTP GET request)
    - Process the registration information at `/u/register` (HTTP POST request)
- Group article related routes under the `/article` route:
    - Serve the article creation page at `/article/create` (HTTP GET request)
    - Process the submitted article at `/article/create` (HTTP POST request)
    - Serve the article page at `/article/view/:article_id ` (HTTP GET request)

## Rendering
- API endpoints and microservices typically respond with data, commonly in JSON format.
- Primarily respond to a request with a HTML template.

## Middleware
- Some common uses of middleware include authorization, validation, etc.

## Installing the Dependency
- `go get -u github.com/gin-gonic/gin`
- Create the `main.go`
- Build the application and create an executable named `app`
    - `go build -o app`
- Execute the application
    - `./app`