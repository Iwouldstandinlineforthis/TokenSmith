#!/bin/bash
rm ./release/* -f
go build -o ./release/tokensmith_linux_x64 main.go
GOOS=windows go build -o ./release/tokensmith_windows_x64.exe main.go
