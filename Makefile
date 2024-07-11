all: build

build:
	go build -o dl

fmt:
	gofmt -s -w .
