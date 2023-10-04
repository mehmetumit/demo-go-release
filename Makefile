COVERAGE_FILE=coverage.txt
BENCHMARKS="Add|Subtract"
all:
build-image:
	@docker build -t demo-go-release:latest . --build-arg="VERSION=$$(git tag -l | tail -n 1)" --build-arg="COMMIT=$$(git rev-parse --short HEAD)"
run-image:
	@docker run --rm demo-go-release:latest
test-covarage:
	go test -v ./cmd -race -coverprofile=${COVERAGE_FILE} -covermode=atomic
	go tool cover -html ${COVERAGE_FILE}
test:
	go test -v ./cmd
test-bench:
	go test ./cmd -bench=${BENCHMARKS}
