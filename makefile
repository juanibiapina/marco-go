MARCOTESTS := $(wildcard ../marco-tests/**/*.mrc)

default: test-marco

test:
	go test ./...

build: test
	go build -o bin/marco cmd/marco/main.go

$(MARCOTESTS):
	@./bin/marco $@ > /dev/null

test-marco: build $(MARCOTESTS)

clean:
	rm -rf bin

.PHONY: build test default clean test-marco $(MARCOTESTS)
