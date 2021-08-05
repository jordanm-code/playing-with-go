#!/bin/bash

# go to work dir
cd /app

# setup our go instance
if [ ! -f go.mod ]; then
  go mod init go-project
fi

# run the go code?
go run .
