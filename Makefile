.PHONY: all test cover
COVER_FILE=./.coverage.txt
BIN=bin/

all: get build test

default: all

get:
	go get ./...

generate: get
	go generate ./...
build: generate
	go build -o $(BIN) ./...

test: generate
	go test ./... -v -coverprofile $(COVER_FILE) && go tool cover -func $(COVER_FILE)

clean:
	$(GOCLEAN)
	rm -f $(BIN) || true &&	rm $(COVER_FILE) || true

install:
	go install ./...

cover: test
	go tool cover -func $(COVER_FILE)
