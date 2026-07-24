VERSION ?= dev

.PHONY: build test fmt vet clean

build:
	mkdir -p bin
	go build -ldflags "-X main.Version=$(VERSION)" -o bin/terraform-provider-hello .

test:
	@bash test/run-test.sh

fmt:
	gofmt -s -w .

vet:
	go vet ./...

clean:
	rm -rf bin dist
