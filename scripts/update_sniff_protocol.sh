#!/usr/bin/env bash

#
# Copyright (c) 2019 Stackinsights to present
# All rights reserved
#

set -e

export BASEDIR=$(dirname "$0")/..
export TEMPDIR="$BASEDIR"/temp
export PROTOCOLDIR="$BASEDIR"/temp/protocols
. "$BASEDIR"/dependencies.sh

function initProtocolHome(){
  if [[ -d "$PROTOCOLDIR" ]]; then
    rm -rf "$PROTOCOLDIR}"
  fi
  mkdir -p "$PROTOCOLDIR"
}

function addProtocol(){
  name=$1
  protocolTarAddress=$2

  curl -sLo "$PROTOCOLDIR"/$name.tgz $2

  if [[ ! -d "$PROTOCOLDIR"/$name ]]; then
    mkdir "$PROTOCOLDIR"/$name
  else
    rm -rf "$PROTOCOLDIR"/$name/*
  fi

  tar -zxf "$PROTOCOLDIR"/$name.tgz -C "$PROTOCOLDIR"/$name --strip 1
}

function cleanHistoryCodes(){
  rm -rf "$BASEDIR"/collect
  rm -rf "$BASEDIR"/proto
  find "$BASEDIR"/satellite -name "*.go" -exec rm {} \;
}

function prepareSatelliteProtocols() {
  mkdir -p "$PROTOCOLDIR/satellite/"
  cp -R "$BASEDIR"/satellite/data/v1/*.proto "$PROTOCOLDIR/satellite/"
  cp -R "$BASEDIR"/satellite/envoy/accesslog/v3/*.proto "$PROTOCOLDIR/satellite/"
}

function generateCodes(){
  go get -u google.golang.org/protobuf/cmd/protoc-gen-go@v1.26.0
  go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1.0

  "$BASEDIR"/scripts/protoc.sh \
    --proto_path="$PROTOCOLDIR"/stackinsights-collect \
    --go_out="$BASEDIR" \
    --go-grpc_out="$BASEDIR" \
    "$PROTOCOLDIR"/stackinsights-collect/*/*.proto \
    "$PROTOCOLDIR"/stackinsights-collect/*/*/*.proto

  "$BASEDIR"/scripts/protoc.sh \
    --experimental_allow_proto3_optional \
    --proto_path="$PROTOCOLDIR"/stackinsights-collect \
    --proto_path="$PROTOCOLDIR"/envoy \
    --proto_path="$PROTOCOLDIR"/xds \
    --proto_path="$PROTOCOLDIR"/protoc-gen-validate \
    --proto_path="$PROTOCOLDIR"/prometheus-model \
    --proto_path="$PROTOCOLDIR"/opentelementry \
    --proto_path="$PROTOCOLDIR" \
    --go_out="$BASEDIR" \
    --go-grpc_out="$BASEDIR" \
    $(bash "$BASEDIR"/scripts/third-proto-import.sh opts "$PROTOCOLDIR") \
    "$PROTOCOLDIR"/satellite/*.proto \
    $(bash "$BASEDIR"/scripts/third-proto-import.sh files "$PROTOCOLDIR")

  mv "$BASEDIR"/demo.stackinsights.ai/repo/goapi/collect "$BASEDIR"/ \
  && mv "$BASEDIR"/demo.stackinsights.ai/repo/goapi/satellite/data/v1/* "$BASEDIR"/satellite/data/v1 \
  && mv "$BASEDIR"/demo.stackinsights.ai/repo/goapi/satellite/envoy/accesslog/v3/* "$BASEDIR"/satellite/envoy/accesslog/v3 \
  && mv "$BASEDIR"/demo.stackinsights.ai/repo/goapi/proto/ "$BASEDIR"/ \
  && rm -rf "$BASEDIR"/stackinsights.apache.org && rm -rf $TEMPDIR

  go mod tidy
}

initProtocolHome
addProtocol stackinsights-collect https://github.com/apache/stackinsights-data-collect-protocol/archive/"${COLLECT_PROTOCOL_SHA}".tar.gz
addProtocol envoy https://github.com/envoyproxy/data-plane-api/archive/${ENVOY_SERVICE_PROTOCOL_SHA}.tar.gz
addProtocol xds https://github.com/cncf/xds/archive/${XDS_SERVICE_PROTOCOL_SHA}.tar.gz
addProtocol protoc-gen-validate https://github.com/envoyproxy/protoc-gen-validate/archive/${PROTOC_VALIDATE_SHA}.tar.gz
addProtocol prometheus-model https://github.com/prometheus/client_model/archive/${PROMETHEUS_MODEL_SHA}.tar.gz
addProtocol opentelementry https://github.com/open-telemetry/opentelemetry-proto/archive/${OPENTELEMETRY_MODEL_SHA}.tar.gz

cleanHistoryCodes
prepareSatelliteProtocols
generateCodes
