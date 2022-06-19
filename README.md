
# GO-Cards API
  
## Available Endpoints

- Create deck
 Creates a deck of cards
- Open deck
 Opens a deck of cards
- Draw Cards
 Draw cards from a deck

The repo includes a postman collection that shows how to call the API and all the parameters.

## How to run

### Using docker

Build the docker image using

```cmd
docker build -f Dockerfile -t go-cards:latest .
```

Run the image using:

```cmd
docker run -p 8080:8080 go-cards
```

### Without docker

Install dependencies

```cmd
go get .
```

Run the app

```cmd
go run main.go
```

## Tests

Tests were added to cover all the available handlers.

 To run the tests, use:

```cmd
go test ./tests
```
