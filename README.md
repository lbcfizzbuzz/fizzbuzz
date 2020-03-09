# fizzbuzz
A simple fizzbuzz REST server written in Go

## Prerequisites
- Go 1.13
- Docker
## Installation
Please first install the database given at: https://github.com/samyy321/fizzbuzz_db.

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
