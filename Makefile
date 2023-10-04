COVERAGE_FILE=coverage.out
BENCHMARKS="Add|Subtract"
all:
build-image:
	@docker build -t demo-go-release:latest . --build-arg="VERSION=$$(git tag -l | tail -n 1)" --build-arg="COMMIT=$$(git rev-parse --short HEAD)"
run-image:
	@docker run --rm demo-go-release:latest
test-covarage:
	go test -v ./cmd -coverprofile=${COVERAGE_FILE}
	go tool cover -html ${COVERAGE_FILE}
test:
	go test -v ./cmd
test-bench:
	go test ./cmd -bench=${BENCHMARKS}
