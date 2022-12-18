.PHONY: all test cover
COVER_FILE=./.coverage.txt
MOCKS=./mocks
BIN=bin/

all: get build test

default: all

get:
	go get ./...

gen: get
	go generate ./...
build: gen
	go build -o $(BIN) ./...

test: gen
	go install github.com/dave/courtney && courtney -o $(COVER_FILE) ./pkg/... && go tool cover -func $(COVER_FILE)

clean:
	$(GOCLEAN)
	rm -rf $(BIN) || true && rm $(COVER_FILE) || true && rm -rf $(MOCKS)

install:
	go install ./...

cover: test
	go tool cover -func $(COVER_FILE)
