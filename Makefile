
-include local.Makefile

.PHONY: lint
lint: ## lint
	golangci-lint run

.PHONY: fmt
fmt: tidy ## tidy,format and imports
	gofumpt -w `find . -type f -name '*.go' -not -path "./vendor/*"`
	goimports -w `find . -type f -name '*.go' -not -path "./vendor/*"`

.PHONY: tidy
tidy: ## go mod tidy
	go mod tidy

build:
	go env
	ls ./cmd/ | xargs -n1 -I {} sh -c "set -ex;CGO_ENABLED=0 go build -v -ldflags '-s -w' -trimpath -o bin/{} ./cmd/{}"
	du -sh bin/*
