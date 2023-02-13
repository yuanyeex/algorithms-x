#!/bin/bash

function cleanTestCache() {
  echo "cleaning..."
  go clean -testcache -v;
}

function test() {
  go test ./...;
}

case $1 in
  clean) cleanTestCache;
  ;;
esac

test;