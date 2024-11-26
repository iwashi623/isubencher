.PHONY: local-init init
init:
	@if [ -z "$(ISUCON_NAME)" ]; then \
	    echo "Error: ISUCON_NAME is not set"; \
	    exit 1; \
	fi
	@echo "Running target: $(ISUCON_NAME)-init"
	mkdir -p modules && rm -rf modules/*
	$(MAKE) $(ISUCON_NAME)-init

local-init: 
	mkdir -p modules && rm -rf modules/*
	$(MAKE) $(ISUCON_NAME)-clone
	$(MAKE) init

.PHONY: kayac-listen80-init kayac-listen80-prepare-data kayac-listen80-build-bench kayac-listen80-clone
kayac-listen80-init: kayac-listen80-clone kayac-listen80-prepare-data kayac-listen80-build-bench
	cp modules/kayac-isucon-2022/bench/bench .
	rm modules/kayac-isucon-2022/bench/bench
	mkdir -p data
	cp -r modules/kayac-isucon-2022/bench/data/*.json data/
	rm modules/kayac-isucon-2022/bench/data/*.json && rm modules/kayac-isucon-2022/sql/90_isucon_listen80_dump.sql

kayac-listen80-prepare-data:
	cd modules/kayac-isucon-2022 && make dataset

kayac-listen80-build-bench:
	cd modules/kayac-isucon-2022/bench && make bench

kayac-listen80-clone:
	cd modules && git clone https://github.com/kayac/kayac-isucon-2022.git
