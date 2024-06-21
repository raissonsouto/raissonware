CMD_DIRS := client server unlocker

# Targets
.PHONY: all build test clean

all: build test

build: $(CMD_DIRS)

$(CMD_DIRS):
	go build -o ./bin/$@ ./cmd/$@/*.go

test:
	go test ./...

clean:
	go clean
	rm -f ./bin/*
