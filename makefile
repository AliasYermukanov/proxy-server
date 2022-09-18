run:
	go run main.go

fmtfix:
	golangci-lint run --fix -E gofmt,gofumpt,goimports

fmt:
	go fmt ./...
