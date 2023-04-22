.PHONY: oapi-codegen
oapi-codegen:
	oapi-codegen -package openapi ./interface/openapi/main.yaml > ./interface/openapi/main.gen.go

.PHONY: ent
ent:
	go generate ./domain/ent

.PHONY: errorcode
errorcode:
	go generate ./server/errorcode

.PHONY: download-module
download-module:
	go mod download

.PHONY: migrate
migrate:
	go run cmd/migration/main.go -env debug
	go run cmd/migration/main.go -env test
