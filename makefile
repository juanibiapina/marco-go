default: test

test:
	go test ./...

build: test
	go build -o bin/marco cmd/marco/main.go

clean:
	rm -rf bin

.PHONY: build test default clean
