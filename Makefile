build:
	go build -o bin/gorderbook

run: build
	./bin/gorderbook

test:
	go test -v ./...