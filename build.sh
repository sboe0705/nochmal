#!/bin/bash

# clean
rm -rf dist

# build
for goos in windows linux; do
  for goarch in amd64 arm64; do
    outfile="nochmal-${goos}-${goarch}"
    [[ "$goos" == "windows" ]] && outfile+=".exe"
    GOOS=$goos GOARCH=$goarch go build -o "dist/$outfile" .
  done
done
