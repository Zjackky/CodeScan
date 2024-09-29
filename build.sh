#!/bin/bash

# Define the list of target operating systems and architectures
os_archs=("darwin:amd64" "darwin:arm64" "linux:amd64" "windows:amd64")

# Define the Go compiler flags
LDFLAGS="-s -w"

# Loop through each OS/architecture pair and build JodeScanner
for pair in "${os_archs[@]}"; do
    os=$(echo "$pair" | cut -d ":" -f 1)
    arch=$(echo "$pair" | cut -d ":" -f 2)
    output="./releases/CodeScan_${os}_${arch}"

    # For Windows, add .exe extension to the output file
    if [[ "$os" == "windows" ]]; then
        output="$output.exe"
    fi

    # Build JodeScanner for the current OS/architecture pair
    echo "Building $output..."
    GOOS="$os" GOARCH="$arch" go build -trimpath -ldflags "$LDFLAGS" -o "$output" main.go
    echo "Build $output done"
done

