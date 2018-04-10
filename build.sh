#!/usr/bin/env sh

# Purpose: This script compiles the Go packages and dependencies.
# Instructions: make build <BINARY>

set -eu

BINARY="$1"

rm -f "$BINARY"
# go get github.com/prometheus/client_golang/prometheus
go build -o ./"$BINARY" ./main.go
