#!/usr/bin/env bash

VERSION=$(cat VERSION)
APP_NAME="internet-outages-monitor"
OUTFILE_LINUX="${APP_NAME}-v${VERSION}-linux_amd64"
OUTFILE_ARM_32="${APP_NAME}-v${VERSION}-linux_arm32"
OUTFILE_ARM_64="${APP_NAME}-v${VERSION}-linux_arm64"
OUTFILE_WINDOWS="${APP_NAME}-v${VERSION}-windows_amd64.exe"
OUTFILE_DARWIN="${APP_NAME}-v${VERSION}-darwin_amd64"

echo "Deleting old builds..."
rm -rf ./dist
mkdir -p ./dist

export CGO_ENABLED=0

echo "Building for linux..."
GOOS=linux GOARCH=amd64 go build -o "dist/${OUTFILE_LINUX}" -ldflags '-w -s -extldflags "-static"' .

echo "Building for arm32..."
GOOS=linux GOARCH=arm GOARM=7 go build -o "dist/${OUTFILE_ARM_32}" -ldflags '-w -s -extldflags "-static"' .

echo "Building for arm64..."
GOOS=linux GOARCH=arm64 go build -o "dist/${OUTFILE_ARM_64}" -ldflags '-w -s -extldflags "-static"' .

echo "Building for windows..."
GOOS=windows GOARCH=amd64 go build -o "dist/${OUTFILE_WINDOWS}" -ldflags '-w -s -extldflags "-static"' .

echo "Building for darwin..."
GOOS=darwin GOARCH=amd64 go build -o "dist/${OUTFILE_DARWIN}" -ldflags '-w -s -extldflags "-static"' .
