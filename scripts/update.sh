#!/usr/bin/env bash

#
# Copyright (c) 2019 Stackinsights to present
# All rights reserved
#

set -e

bash "$(dirname "$0")"/update_sniff_protocol.sh
bash "$(dirname "$0")"/update_query_protocol.sh

rm -rf temp
