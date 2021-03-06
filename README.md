# fizzbuzz
A simple fizzbuzz REST server written in Go

## Prerequisites
- Go 1.13 (https://golang.org/doc/install)
- Docker (https://docs.docker.com/install/)
## Installation
Please first install the database given at: https://github.com/lbcfizzbuzz/fizzbuzz_db.

Go to `cmd/fizzbuzz` and build the project:

```
$ cd ./cmd/fizzbuzz
$ go build
```
## Usage
Run the fizzbuzz executable by using the default configuration (`config/config-dev.json`):
```
$ cd ./fizzbuzz
```
Run the fizzbuzz executable and pass a configuration file to it (must be formatted as explained in `config/README.md`):
```
$ cd ./fizzbuzz -config "path/to/configuration/file"
```

The api accepts requests at the `http://localhost:8080/` url.

The `fizzbuzz` endpoint accepts five parameters, for example `http://localhost:8080/fizzbuzz?int1=3&int2=5&limit=10&str1=fizz&str2=buzz` will do a fizzbuzz from 1 to 10 included.

The `statistics` endpoint accepts no parameters, for example `http://localhost:8080/statistics`.

## Tests
To run the tests just type:
```
$  go test github.com/lbcfizzbuzz/fizzbuzz/tests
```
## Project structure
- `cmd/fizzbuzz` contains the main
- `config`contains the server configuration files and the code to read them
- `core` contains the fizzbuzz algorithm
- `datastore` contains the code for the communication with a database, there is an interface that can be implemented
to add a new type of storage
- `internal/constants` contains the constants of the project
- `models`contains structures that represents our business objects
- `server` contains the code to handle the clients requests
- `service` contains the code that does the link between the server and the datastore
- `tests` contains the tests
## What could be improved
- Create a separate table for request parameters to allow adding or removing parameters easily

Feel free to send me you feedbacks by creating a Github issue.
