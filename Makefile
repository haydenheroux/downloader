all: build

build:
	go build -o bin/ ./cmd/media

fmt:
	gofmt -s -w .

install:
	go install
