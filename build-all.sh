#!/bin/bash

# First we update the version
if [ $# -ne 1 ]; then
    echo "Usage: $0 <semantic_version>"
    exit 1
fi

version="$1"
content="package constants

const (
    Name    = \"gov\"
    Version = \"$version\"
)
"

echo "$content" > ./constants/base.go
echo "$version" > ./version

# Then we build the binaries
supported_combinations=($(go tool dist list | grep -vE 'android|ios' | awk -F/ '{print $1 "-" $2}'))
app_file="main.go"
output_dir="bin"

for combination in "${supported_combinations[@]}"; do
  # Split the combination into platform and architecture
  platform="${combination%-*}"
  arch="${combination#*-}"

  export GOOS="$platform"
  export GOARCH="$arch"

  output_name="$output_dir/$platform/$arch/gov"

  go build -o "$output_name" "$app_file"

  unset GOOS
  unset GOARCH

  # Check if it's Windows and add the .exe extension
  if [ "$platform" == "windows" ]; then
    mv "$output_name" "$output_name.exe"
  fi
done

echo "Cross-compilation complete. Binaries are in the '$output_dir' directory."
