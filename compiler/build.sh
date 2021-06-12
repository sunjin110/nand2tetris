#!/bin/sh

echo "=== build compiler ==="
go build -o compiler ./cmd/compiler/main.go

echo "=== build analyzer ==="
go build -o analyzer ./cmd/analyzer/main.go

