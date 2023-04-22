MAKEFILE_DIR := $(dir $(lastword $(MAKEFILE_LIST)))

.PHONY: install-tools
install-tools:
	go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@v1.11.0
	go install entgo.io/ent/cmd/ent@v0.10.1

.PHONY: download-module
download-module:
	make -C package download-module
	make -C main download-module

.PHONY: generate-code
generate-code:
	make -C main oapi-codegen
	make -C main ent
	make errorcode

.PHONY: migrate
migrate:
	make -C main migrate

.PHONY: errorcode
errorcode:
	cd main
	make -C main errorcode
	cd $(MAKEFILE_DIR)

.PHONY: setup
setup:
	make install-tools
	make download-module
	make generate-code
	make migrate