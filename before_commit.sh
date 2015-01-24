pushd `dirname "${0}"`
  echo "go fmt"
  pushd tumblr
    find *.go -type f | go fmt
  popd
popd
