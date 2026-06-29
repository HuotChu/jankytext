BINARY := jankytext
VERSION ?= 0.1.2

.PHONY: test build install dist clean

test:
	go test ./...

build:
	mkdir -p bin
	go build -o bin/$(BINARY) ./cmd/jankytext

install:
	go install ./cmd/jankytext

dist:
	rm -rf dist
	mkdir -p dist
	GOOS=darwin GOARCH=arm64 go build -o dist/$(BINARY) ./cmd/jankytext
	tar -C dist -czf dist/$(BINARY)_$(VERSION)_darwin_arm64.tar.gz $(BINARY)
	rm dist/$(BINARY)
	GOOS=darwin GOARCH=amd64 go build -o dist/$(BINARY) ./cmd/jankytext
	tar -C dist -czf dist/$(BINARY)_$(VERSION)_darwin_amd64.tar.gz $(BINARY)
	rm dist/$(BINARY)
	GOOS=linux GOARCH=arm64 go build -o dist/$(BINARY) ./cmd/jankytext
	tar -C dist -czf dist/$(BINARY)_$(VERSION)_linux_arm64.tar.gz $(BINARY)
	rm dist/$(BINARY)
	GOOS=linux GOARCH=amd64 go build -o dist/$(BINARY) ./cmd/jankytext
	tar -C dist -czf dist/$(BINARY)_$(VERSION)_linux_amd64.tar.gz $(BINARY)
	rm dist/$(BINARY)
	GOOS=windows GOARCH=amd64 go build -o dist/$(BINARY).exe ./cmd/jankytext
	cd dist && zip -q $(BINARY)_$(VERSION)_windows_amd64.zip $(BINARY).exe
	rm dist/$(BINARY).exe
	GOOS=windows GOARCH=arm64 go build -o dist/$(BINARY).exe ./cmd/jankytext
	cd dist && zip -q $(BINARY)_$(VERSION)_windows_arm64.zip $(BINARY).exe
	rm dist/$(BINARY).exe
	cd dist && shasum -a 256 * > checksums.txt

clean:
	rm -rf bin dist
