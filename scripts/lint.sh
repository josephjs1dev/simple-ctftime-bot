#!/usr/bin/env bash

if ! command -v ./bin/golangci-lint > /dev/null; then
  echo "Installing golangci-lint"
  curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s latest
fi

./bin/golangci-lint run $*
