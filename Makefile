# Build the Go binary
build:
	go build -o campaingdemo ./cmd

# Build the Docker image
docker-build:
	docker build -t campaingdemo .

# Run the Docker container
docker-run:
	docker run --name campaingdemo-app campaingdemo
# Clean up the Go binary
clean:
	rm -f campaingdemo

# Clean up the Docker image
docker-clean:
	docker stop campaingdemo-app || true
	docker rm campaingdemo-app || true
	docker rmi campaingdemo
# Run tests
test:
	go test ./...

.PHONY: build docker-build docker-run clean docker-clean test
