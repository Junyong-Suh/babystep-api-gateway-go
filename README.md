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
Hello
```
We will start from here :)

## API documentation

### Root
Returns a string "Hello"
```
$ curl localhost:8080
Hello
```

### /health
Returns a service name and version info in JSON
```
$ curl localhost:8080 | jq
{
  "version": "v0.0.1",
  "service": "babystep-api-gateway"
}
```

### /echo
Returns a response from downstream (echo-api.3scale.net)
```
$ curl localhost:8080/echo | jq
{
  "method": "GET",
  "path": "/",
  "args": "",
  "body": "",
  "headers": {
    "HTTP_VERSION": "HTTP/1.1",
    "HTTP_HOST": "echo-api.3scale.net",
    "HTTP_ACCEPT_ENCODING": "gzip",
    "HTTP_USER_AGENT": "Go-http-client/1.1",
    "HTTP_X_FORWARDED_FOR": "218.39.203.11, 10.0.103.17",
    "HTTP_X_FORWARDED_HOST": "echo-api.3scale.net",
    "HTTP_X_FORWARDED_PORT": "443",
    "HTTP_X_FORWARDED_PROTO": "https",
    "HTTP_FORWARDED": "for=10.0.103.17;host=echo-api.3scale.net;proto=https"
  },
  "uuid": "87561cda-aefe-4062-af81-6307551c9842"
}
```

## Tests
As we keep all tests under ./tests folder, run following command to trigger unit tests
```
$go test ./tests
```
