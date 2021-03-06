#!/usr/bin/env bash
echo "" > coverage.tmp
echo 'mode: atomic' > coverage.txt && go list ./... | grep -v vendor | xargs -n1 -I{} sh -c 'go test -covermode=atomic -coverprofile=coverage.tmp {} && tail -n +2 coverage.tmp >> coverage.txt' && rm coverage.tmp