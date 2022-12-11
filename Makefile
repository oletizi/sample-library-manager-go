.PHONY: all test cover
COVER_FILE=./.coverage.txt
BIN=bin/

all: get build test

default: all

get:
	go get ./...

build:
	go build -o $(BIN) ./...

test:
	go test ./... -v -coverprofile $(COVER_FILE) && go tool cover -func $(COVER_FILE)

clean:
	$(GOCLEAN)
	rm -f $(BIN) || true &&	rm $(COVER_FILE) || true

install:
	go install ./...

cover: test
	go tool cover -func $(COVER_FILE)
