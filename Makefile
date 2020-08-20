all: clean build compile
.PHONY: all

.PHONY: clean
clean:
	rm -rf ./bin

.PHONY: build
build:
	go build -o bin/opendentalgo opendentalgo.go

.PHONY: run
run:
	go run opendentalgo.go

.PHONY: compile
compile:
	GOOS=linux GOARCH=amd64 go build -o bin/main-linux-amd64 opendentalgo.go
