# go-exercise

## Running the app

```bash
$ go run cmd/server/main.go
```

## Tests

Install [mockery](https://github.com/vektra/mockery):

```bash
$ go install github.com/vektra/mockery/v2@v2.43.2
```

Generate mocks and run all tests:

```bash
$ go generate ./...
$ go test -v ./...
```
