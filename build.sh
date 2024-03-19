#!/bin/bash
set -euo pipefail

export GOOS=linux
for exe in detect build; do
  go build -ldflags="-s -w" -o ./bin/${exe} ./cmd/${exe}/main.go
done
