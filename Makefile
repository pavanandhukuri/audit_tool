.PHONY: test

# Run Tests
test:
	@echo "Running tests..."
	@go test -v ./...
	@echo "Tests passed."

# Run tests with coverage
test-coverage:
	@echo "Running tests with coverage..."
	@go test -coverprofile=coverage.out ./...
	@go tool cover -html=coverage.out
	@rm coverage.out
	@echo "Tests passed."

# Generate mock files using mockery
mock:
	@echo "Generating mocks..."
	@mockery --all --case underscore
	@echo "Mocks generated."