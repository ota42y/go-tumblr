#!/bin/bash

echo "go fmt example/*"
go fmt example/*

pushd tumblr
echo "go fmt"
find . -name "*.go" -exec go fmt {} \;

echo "golint"
golint
popd
