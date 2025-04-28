#!/usr/bin/env bash
set -o errexit
set -o nounset
set -o xtrace

GOLANGCI_LINT_RUN_OPTS=""
GOLANGCI_LINT_RUN_OPTS="$GOLANGCI_LINT_RUN_OPTS --verbose" # debug linter timing and mem usage
GOLANGCI_LINT_RUN_OPTS="$GOLANGCI_LINT_RUN_OPTS --concurrency=2" # prow job lags if too many threads
export GOLANGCI_LINT_RUN_OPTS

make lint
