.PHONY: test
## test: test the application
test:
	go test ./...

.PHONY: build
## build: build the application
build:
	sh ./build.sh

.PHONY: help
## help: prints this help message
help:
	@echo "Usage: "
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'