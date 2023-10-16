# Build image and run container
.PHONY: up
up: build run
	@echo "This is CS741 for rail cipher test"
	@echo "Use 'docker exec rail <cmd> to run any commands"
	@echo "  Example: docker exec rail go test -v"

# Close the running container and remove it from background
.PHONY: down
down:
	@docker stop rail && docker container rm rail && echo "Container rail ended."

# Build rail image
.PHONY: build
build:
	@docker build -t rail:latest -f rail.dockerfile .

# Start a running container from rail image
.PHONY: run
run:
	@docker run -d --name rail rail:latest

# Run test in container
.PHONY: test
test:
	@docker exec rail go test -v

