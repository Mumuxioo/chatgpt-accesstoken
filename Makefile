GOPATH:=$(shell go env GOPATH)

.PHONY: init
	@go mod tidy

snapshots:
	@goreleaser check
	@goreleaser release --snapshot --skip-publish --rm-dist

docker:
	@docker build -t askaigo/chatgpt-accesstoken:v1.0.0 .