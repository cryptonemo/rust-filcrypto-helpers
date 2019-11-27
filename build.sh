#!/bin/bash

set -e

staging="./tmp_staging"

echo "Cleaning build area ..."
cargo clean
rm -rf $staging
rm -f ./bindings.go

# Setup temporary staging dir
mkdir -p $staging
pushd $staging > /dev/null

echo "Building project and gathering linker flags ..."
RUSTFLAGS='--print native-static-libs' cargo build --release --all > build_output 2>&1
linker_flags=$(cat build_output | grep native-static-libs | head -n 1 | cut -d ':' -f 3)

popd > /dev/null
rm -rf $staging

echo "Updating binding linker flags ..."
cp ./bindings.go.template ./bindings.go
sed -i "s/@LIBS@/${linker_flags}/g" bindings.go

echo "Verifying that bindings work"
go build ./bindings.go

echo "Build complete!"
echo "Run tests using: go test ./bindings_test.go -v"
