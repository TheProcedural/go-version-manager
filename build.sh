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

supported=("linux/amd64" "linux/arm64" "linux/386" "linux/arm" "darwin/amd64" "darwin/arm64")
app_file="main.go"
output_dir="bin"

for combo in "${supported[@]}"; do
  # Split the combo into platform and architecture
  IFS='/' read -r -a parts <<< "$combo"
  platform="${parts[0]}"
  arch="${parts[1]}"

  export GOOS="$platform"
  export GOARCH="$arch"

  output_name="$output_dir/$platform/$arch/gov"

  go build -o "$output_name" "$app_file"

  unset GOOS
  unset GOARCH
done

echo "Cross-compilation complete. Binaries are in the '$output_dir' directory."
