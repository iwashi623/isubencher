PHONY: init

init: prepare-data build-bench

prepare-data:
	cd kayac-isucon-2022 && make dataset

build-bench:
	cd kayac-isucon-2022/bench && make bench
