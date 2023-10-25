#!/usr/bin/env bash

#
# Copyright (c) 2019 Stackinsights to present
# All rights reserved
#

set -e

BASEDIR=$(dirname "$0")/..
TEMPDIR="$BASEDIR"/temp

. "$BASEDIR"/dependencies.sh

if [[ ! -d "$TEMPDIR" ]]; then
  mkdir -p "$TEMPDIR"
else
  rm -rf "${TEMPDIR:?}"/*
fi

curl -sLo "$TEMPDIR"/query-protocol.tgz https://github.com/apache/stackinsights-query-protocol/archive/"${QUERY_PROTOCOL_SHA}".tar.gz

if [[ ! -d "$TEMPDIR"/query-protocol ]]; then
  mkdir "$TEMPDIR"/query-protocol
else
  rm -rf "$TEMPDIR"/query-protocol/*
fi

tar -zxf "$TEMPDIR"/query-protocol.tgz -C "$TEMPDIR"/query-protocol --strip 1

rm -rf "$TEMPDIR"/query-protocol.tgz

go get github.com/99designs/gqlgen@v0.17.23

"$(go env GOPATH)"/bin/gqlgen -h > /dev/null 2>&1 || GO111MODULE=off go get github.com/99designs/gqlgen
go run "$BASEDIR"/scripts/tools/query_mutation.go

rm -rf "$TEMPDIR"/query-protocol

go mod tidy
