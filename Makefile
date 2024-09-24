.PHONY: test-coverage

test-coverage:
	mkdir -p coverage
	-go test -coverprofile=coverage/coverage.out ./...
	go tool cover -html=coverage/coverage.out -o coverage/coverage.html