#!/usr/bin/env bash
set -o errexit
set -o nounset
set -o xtrace

export GOMAXPROCS=2 # prow job lags if too many threads

make test
