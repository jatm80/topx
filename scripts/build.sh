#!/bin/bash

set -euo pipefail 

cd src

go mod vendor
go test -v ./...
go build -o ../bin/topx
