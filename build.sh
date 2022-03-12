#!/usr/bin/env bash

VERSION=$(cat VERSION)
APP_NAME="internet-outages-monitor"
OUTFILE_LINUX="${APP_NAME}-v${VERSION}-linux_amd64"
OUTFILE_ARM="${APP_NAME}-v${VERSION}-linux_arm32"
OUTFILE_WINDOWS="${APP_NAME}-v${VERSION}-windows_amd64.exe"

echo "Deleting old builds..."
rm -rf ./dist
mkdir -p ./dist

export CGO_ENABLED=0

echo "Building for linux..."
GOOS=linux GOARCH=amd64 go build -o "dist/${OUTFILE_LINUX}" -ldflags '-w -s -extldflags "-static"' .

echo "Building for arm..."
GOOS=linux GOARCH=arm GOARM=7 go build -o "dist/${OUTFILE_ARM}" -ldflags '-w -s -extldflags "-static"' .

echo "Building for windows..."
GOOS=windows GOARCH=amd64 go build -o "dist/${OUTFILE_WINDOWS}" -ldflags '-w -s -extldflags "-static"' .
