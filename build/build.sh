#!/bin/sh

set -o errexit
set -o nounset
set -o pipefail

# Check if OS, architecture and application version variables are set in Makefile
if [ -z "${OS:-}" ]; then
    echo "OS must be set"
    exit 1
fi
if [ -z "${ARCH:-}" ]; then
    echo "ARCH must be set"
    exit 1
fi
if [ -z "${VERSION:-}" ]; then
    echo "VERSION must be set"
    exit 1
fi

# Disable C code, enable Go modules.
# TODO. Need fix the docker run error -
# 'standard_init_linux.go:211: exec user process caused "exec format error"'
# It happens OS=darwin. So hard code to linux as a workaround
export CGO_ENABLED=0
export GOARCH="${ARCH}"
# export GOOS="${OS}"
export GOOS="linux"
export GO111MODULE=on
export GOFLAGS="-mod=vendor"

# Build the application. `-ldflags -X` sets version variable in importpath for each `go tool link` invocation
go install                                                      \
    -installsuffix "static"                                     \
    -ldflags "-X $(go list -m)/pkg/version.VERSION=${VERSION}"  \
    ./...