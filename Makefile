.PHONY: set-up
set-up:
	mkdir -p modules && rm -rf modules/*

.PHONY: init
init: set-up
	$(MAKE) $(ISUCON_NAME)-init

# ------------------kayac-listen80------------------
.PHONY: kayac-listen80-clone
kayac-listen80-clone: set-up
	cd modules && git clone https://github.com/kayac/kayac-isucon-2022.git

.PHONY: kayac-listen80-data
kayac-listen80-data: kayac-listen80-clone
	cd modules/kayac-isucon-2022 && make dataset

.PHONY: kayac-listen80-build-bench
kayac-listen80-build-bench: kayac-listen80-data
	cd modules/kayac-isucon-2022/bench && make bench

.PHONY: kayac-listen80-init
kayac-listen80-init: kayac-listen80-build-bench
	cp modules/kayac-isucon-2022/bench/bench .
	rm modules/kayac-isucon-2022/bench/bench
	mkdir -p data
	cp -r modules/kayac-isucon-2022/bench/data/*.json data/
	rm modules/kayac-isucon-2022/bench/data/*.json && rm modules/kayac-isucon-2022/sql/90_isucon_listen80_dump.sql
