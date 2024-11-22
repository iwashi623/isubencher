PHONY: init prepare-data build-bench

init: prepare-data build-bench
	cp kayac-isucon-2022/bench/bench .
	rm kayac-isucon-2022/bench/bench
	mkdir -p data
	cp -r kayac-isucon-2022/bench/data/*.json data/
	rm kayac-isucon-2022/bench/data/*.json && rm kayac-isucon-2022/sql/90_isucon_listen80_dump.sql

prepare-data:
	cd kayac-isucon-2022 && make dataset

build-bench:
	cd kayac-isucon-2022/bench && make bench
