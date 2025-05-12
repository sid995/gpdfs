build:
	@echo "Removing old binaries..."
	@rm -rf bin
	@echo "Building the project..."
	@go build -o bin/fs
	@echo "Providing permissions"
	@chmod +x bin/fs

run: build
	@./bin/fs

test:
	@echo "Running tests..."
	go test -v ./...
	@echo "All tests passed."
	