BINARY := jankytext

.PHONY: test build install clean

test:
	go test ./...

build:
	mkdir -p bin
	go build -o bin/$(BINARY) ./cmd/jankytext

install:
	go install ./cmd/jankytext

clean:
	rm -rf bin dist
