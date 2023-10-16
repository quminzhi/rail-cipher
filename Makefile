# Build image and run container
.PHONY: up
up: build run

# Close the running container and remove it from background
.PHONY: down
down: clean
	@docker stop rail && docker container rm rail && echo "Container rail ended."

# Build rail image
.PHONY: build
build:
	@cd ./main && go build -o ../cipher .
	@docker build -t rail:latest -f rail.dockerfile .
	@docker run -d --name rail rail:latest

# Start a running container from rail image
.PHONY: run
run:
	@docker exec -it rail /bin/bash

# Run test in container
.PHONY: test
test:
	@cd rail && go test -v .

.PHONY: clean
clean:
	@rm -rf cipher

