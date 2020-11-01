#!/usr/bin/env bash

if ! command -v mockery > /dev/null; then
  go get github.com/vektra/mockery/v2/.../
fi

gen_go() {
  echo "Running go generate"
  go generate ./...
}

case "$1" in
go)
  gen_go
  exit
  ;;
*)
  exit
  ;;
esac
