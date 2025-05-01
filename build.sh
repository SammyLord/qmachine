#!/bin/bash
rm -rf build
# Create build directory if it doesn't exist
mkdir -p build

# Function to build for a specific platform
build() {
    local GOOS=$1
    local GOARCH=$2
    local output_name=$3
    
    echo "Building for $GOOS/$GOARCH..."
    GOOS=$GOOS GOARCH=$GOARCH go build -o "build/qmachine_${output_name}"
}

# Build for different platforms
build "darwin" "amd64" "darwin_amd64"
build "darwin" "arm64" "darwin_arm64"
build "linux" "amd64" "linux_amd64"
build "linux" "arm64" "linux_arm64"
build "windows" "amd64" "windows_amd64.exe"

echo "Build complete! Check the build directory for the binaries." 