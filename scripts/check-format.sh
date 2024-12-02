#!/usr/bin/env sh

cd "${GITHUB_WORKSPACE}"

GOFILES="$(find . -name '*.go')"
OUTPUT="$(gofmt -l -d ${GOFILES})"
if [ -n "${OUTPUT}" ]; then
    echo "Bad Format:"
    echo "${OUTPUT}"
    exit 1
fi