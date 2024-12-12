#!/bin/bash -e

go mod tidy
go build -o ./sec ./gateway/main.go
