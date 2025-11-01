build:
	@go build -o bin/bankgoals

run: build
	@./bin/bankgoals

test:
	@go test -v ./..