#!/usr/bin/env bash

set -e

cd "$(dirname "$0")"
find . -name "mock_*.go" -delete
mockery --quiet
