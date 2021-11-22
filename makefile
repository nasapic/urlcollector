DEV_TAG=dev
STG_TAG=stage
PROD_TAG=v0.0.1
IMAGE_NAME=nasapic/urlcollector

# Misc
BINARY_NAME=urlcollector
BINARY_UNIX=$(BINARY_NAME)_unix

build:
	go build -o ./bin/$(BINARY_NAME) ./cmd/$(BINARY_NAME).go

format:
	gofmt -s -w .

run:
	make build
	make setsampleapikey
	./bin/urlcollector -json-api-port 8080 -nasa-api-key-envar API_KEY -logging-level all

setsampleapikey:
	export API_KEY="DEMO_KEY"

