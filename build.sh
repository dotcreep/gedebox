#!/bin/bash

APP="gedebox"

declare -A platforms=(
    ["linux/amd64"]="$APP-linux-amd64"
    ["linux/386"]="$APP-linux-386"
    ["linux/arm"]="$APP-linux-arm"
    ["linux/arm64"]="$APP-linux-arm64"
    ["darwin/amd64"]="$APP-darwin-amd64"
    ["darwin/arm64"]="$APP-darwin-arm64"
    ["windows/amd64"]="$APP-windows-amd64.exe"
    ["windows/arm64"]="$APP-windows-arm64.exe"
)

if [[ ! -d "bin" ]]; then
    mkdir bin
fi

for platform in "${!platforms[@]}"; do
    IFS="/" read -r os arch <<< "$platform"
    output=${platforms[$platform]}
    echo "Build for $os $arch..."
    GOOS=$os GOARCH=$arch go build -o ./bin/$output cmd/main.go
done

echo "Build done!"

