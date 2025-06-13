#!/bin/bash
if [[ "$1" == "debug" ]]; then
  LOG_LEVEL=debug
fi
LOG_LEVEL=$LOG_LEVEL go test -v ./...
