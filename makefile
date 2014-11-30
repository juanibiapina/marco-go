default: test-marco

test:
	go test ./...

build: test
	go build -o bin/marco cmd/marco/main.go

test-marco: build
	./bin/marco ../marco-tests/syntax/list.mrc

clean:
	rm -rf bin

.PHONY: build test default clean test-marco
