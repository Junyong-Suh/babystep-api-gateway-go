#!/bin/bash

echo -e "\033[0;36m Testing: go test ./... \033[0;97m"
go test ./...
echo -e "\033[0;32m Building: go build . \033[0;97m"
go build .
echo -e "\033[0;94m Running: go run . \033[0;97m"
go run .
