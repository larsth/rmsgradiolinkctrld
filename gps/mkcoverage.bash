#!/bin/bash

rm -f ./coverage.out
go test -coverprofile=coverage.out 
go tool cover -html=./coverage.out -o ./coverage.html
