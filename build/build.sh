#!/bin/bash

# Exit on any error
set -e

# Specify the go file to build
GOFILE="../main.go"

# Use the first command-line argument as the output name, or default to "saas-squash"
OUTNAME="${1:-saas-squash}"

# Check if the "small" argument is passed
LDFLAGS=""
if [ "$2" == "small" ]; then
  LDFLAGS="-ldflags \"-s -w\""
fi

# Iterate through OS and ARCH
for GOOS in darwin linux windows; do
  for GOARCH in amd64; do
    # Export the environment variables
    export GOOS=$GOOS
    export GOARCH=$GOARCH
    
    # Name of the output binary file
    OUTFILE="${OUTNAME}-$GOOS-$GOARCH"
    
    # If it's a Windows build, add the .exe extension
    if [ "$GOOS" == "windows" ]; then
      OUTFILE="$OUTFILE.exe"
    fi
    
    # Run the build command
    echo "Building $OUTFILE"
    eval "go build $LDFLAGS -o $OUTFILE $GOFILE"
  done
done

echo "Build process completed"