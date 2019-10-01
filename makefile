.PHONY: test

lint:
	go fmt ./...
	go run "golang.org/x/tools/cmd/goimports" -l -w .
	go run "github.com/golangci/golangci-lint/cmd/golangci-lint" run --enable-all --disable gochecknoinits --disable gochecknoglobals
test:lint
	go test  -cover ./... -v

run:lint
	go run main.go