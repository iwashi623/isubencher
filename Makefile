.PHONY: init
init:
	@if [ -z "$(COMPETITION)" ]; then \
	    echo "Error: COMPETITION is not set"; \
	    exit 1; \
	fi
	@echo "Running target: $(COMPETITION)-init"
	$(MAKE) $(COMPETITION)-init

.PHONY: kayac-listen80-init kayac-listen80-prepare-data kayac-listen80-build-bench
kayac-listen80-init: kayac-listen80-prepare-data kayac-listen80-build-bench
	cp kayac-isucon-2022/bench/bench .
	rm kayac-isucon-2022/bench/bench
	mkdir -p data
	cp -r kayac-isucon-2022/bench/data/*.json data/
	rm kayac-isucon-2022/bench/data/*.json && rm kayac-isucon-2022/sql/90_isucon_listen80_dump.sql

kayac-listen80-prepare-data:
	cd kayac-isucon-2022 && make dataset

kayac-listen80-build-bench:
	cd kayac-isucon-2022/bench && make bench
