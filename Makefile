all:
build-image:
	@docker build -t demo-go-release:latest . --build-arg="VERSION=$$(git tag -l | tail -n 1)" --build-arg="COMMIT=$$(git rev-parse --short HEAD)"
run-image:
	@docker run --rm demo-go-release:latest
