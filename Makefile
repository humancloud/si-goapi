#
# Copyright (c) 2019 Stackinsights to present
# All rights reserved
#

default: update

.PHONY: update-sniff
update-sniff:
	bash scripts/update_sniff_protocol.sh

.PHONY: update-query
update-query:
	bash scripts/update_query_protocol.sh

.PHONY: update
update:
	bash scripts/update.sh

.PHONY: check
check:
	go mod tidy > /dev/null
	@if [ ! -z "`git status -s`" ]; then \
		echo "Following files are not consistent with CI:"; \
		git status -s; \
		git diff; \
		exit 1; \
	fi
