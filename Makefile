# Run the suit (unit) tests
.PHONY: test
test:
	go test ./tests/ -v


# Create the container image
.PHONY: build
build:
	docker build -t devops-api-initial .