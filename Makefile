mocks:
	go mod tidy
	CGO_ENABLED=0 go install github.com/vektra/mockery/v2@latest
	mockery -all
