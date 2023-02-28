.PHONY: build

build:
	go build -o ./bin/app/findtty ./cmd/findtty/main.go

build-integration-setup:
	go build -o ./bin/integration-test/setup ./cmd/integration-test-setup/main.go