run:
	go run main.go

fmt:
	go fmt ./...

check:
	golangci-lint run