all: clean build
.PHONY: all

.PHONY: clean
clean:
	rm -rf ./bin

.PHONY: build
build:
	GOOS=linux GOARCH=amd64 go build -o bin/main-linux-amd64 github.com/jkeam/opendentalgo
