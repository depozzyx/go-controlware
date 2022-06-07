#!/bin/bash     

env GOOS=windows GOARCH=386 go build -o build/i386.exe
env GOOS=windows GOARCH=amd64 go build -o build/amd64.exe
go build -o build/ubuntu

# echo (cat shared/models.go | grep "Version" | grep -o "\".*\"") (date) > build/info.txt