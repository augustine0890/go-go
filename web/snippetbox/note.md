# Web Development

## Project Structure and Organization
- The `cmd` directory will contain the application specific code for the executable apps in project. The web application which will live under the `cmd/web` directory.
- The `pkg` directory will contain the non-application code used in the project. It holds potentially reusable code like validation helpers and the SQL database models for the project.
- The `ui` directory will contain the user-interface assets used by the web app.

## How to run the app
- `go run ./cmd/web`
- Using the `-addr` flag
  - `go run ./cmd/web -addr=":3000"`