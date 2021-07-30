# Go Swagger
- Generate the specification in JSON format
  - `swagger generate spec –o ./swagger.json`
- Load the generated spec in the Swagger UI
  - `swagger serve ./swagger.json`
- Set the flavor flag to `swagger`
  - `swagger serve -F swagger ./swagger.json`