COMMIT := `git rev-parse --short HEAD`
BRANCH := `git symbolic-ref -q --short HEAD`
BUILD_TIME := `date +%FT%T%z`

LDFLAGS += -X=main.GitCommit=$(COMMIT)
LDFLAGS += -X=main.GitBranch=$(BRANCH)
LDFLAGS += -X=main.BuildTime=$(BUILD_TIME)


GO_SRC = $(shell find . -type f -name '*.go' -not -path "./vendor/*")


.PHONY: build
build: $(GO_SRC)
	@go build -ldflags "${LDFLAGS}" -o build/service cmd/service/main.go
