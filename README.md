# babystep-api-gateway-go
Babystep API gateway is a simple API gateway in Go

## Sample
Sample code is from https://github.com/gorilla/mux#full-example
```
$ go get -u github.com/gorilla/mux
$ go build sample.go
$ go run sample
```
'Hello world'-like server is running on localhost:8000
```
$ curl localhost:8000
{
  "serviceName": "babystep-api-gateway",
  "version": "v0.0.1"
}
```
We will start from here :)

## Tests
As we keep all tests under ./tests folder, run following command to trigger unit tests
```
$go test ./tests
```
