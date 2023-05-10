#!/bin/bash

echo "exec 'go mod tidy -compat=1.20'"
go mod tidy -compat=1.20

echo "exec 'go build -o ethsuper'"
go build -o ethsuper
