.DEFAULT_GOAL := build

fmt:
		go fmt ./...

govulncheck: fmt
		go run golang.org/x/vuln/cmd/govulncheck@latest ./...

build: govulncheck
		CGO_ENABLED=0 go build -o ./bin