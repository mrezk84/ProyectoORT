#!/usr/bin/env bash
# Stops the process if something fails
set -xe

# All of the dependencies needed/fetched for your project. 
# FOR EXAMPLE:
go get "github.com/labstack/echo/v4"

# create the application binary that eb uses
GOOS=linux GOARCH=amd64 go build -o bin/application -ldflags="-s -w"
