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
Run the fizzbuzz executable, it will use the configuration file `config/config-dev.json` by default.

You can send requests to the api using the `http://localhost:8080/` url.

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
- `models`contains structures that represents our business objects
- `server` contains the code to handle the requests to the api
- `service` contains the code that does the link between the server and the datastore (could be called controller)
- `tests` contains the tests
## What could be improved
- Adding some logs
- Not making blocking calls to the database
- Create a separate table for request parameters to allow adding or removing parameters easily

Feel free to send me you feedbacks by creating a Github issue.
