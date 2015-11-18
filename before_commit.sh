#!/bin/bash

echo "go fmt"
find . -name "*.go" -exec go fmt {} \;

pushd tumblr

echo "golint"
golint

popd
