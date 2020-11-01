#!/usr/bin/env bash

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
