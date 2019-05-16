#!/bin/bash

echo -e "\033[0;32m Building: go build main.go \033[0;97m"
go build main.go
echo -e "\033[0;36m Testing: go test ./... \033[0;97m"
go test ./...
echo -e "\033[0;94m Running: go run main.go \033[0;97m"
go run main.go
