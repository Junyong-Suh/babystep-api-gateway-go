# Babaystep API Gateway
Babaystep API Gateway is a simple API gateway in Go

## Dependencies
If you add any dependencies, please add in go.mod then
```
$ go mod init
$ go mod tidy
```

## Run
```
$ go run .
```

## Test
```
$ go test ./...
```

## API documentation

### /
Returns a string "Hello"
```
$ curl localhost:8080
Hello
```

### /health
Returns a service name and version info in JSON
```
$ curl localhost:8080/health | jq
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
