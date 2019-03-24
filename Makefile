COMMIT := `git rev-parse --short HEAD`
BRANCH := `git symbolic-ref -q --short HEAD`
BUILD_TIME := `date +%FT%T%z`

LDFLAGS += -X=main.GitCommit=$(COMMIT)
LDFLAGS += -X=main.GitBranch=$(BRANCH)
LDFLAGS += -X=main.BuildTime=$(BUILD_TIME)

GO_SRC = $(shell find . -type f -name '*.go' -not -path "./vendor/*")

SERVICE_NAME="greeter"

DOCKER_IMAGE="lukasjarosch/microservice-structure"
DOCKER_TAG="v-${COMMIT}"

.PHONY: build
build: $(GO_SRC)
	@go build -ldflags "${LDFLAGS}" -o build/service cmd/service/main.go

.PHONY: run
run:
	@go run -ldflags "${LDFLAGS}" cmd/service/main.go

.PHONY: docker
docker:
	@echo "--> building docker image"
	docker build . -t ${DOCKER_IMAGE}:${DOCKER_TAG} --build-arg GO_VERSION=1.11 -f ./Dockerfile

.PHONY: docker-run
docker-run:
	@echo "--> starting container ${DOCKER_IMAGE}:${DOCKER_TAG}"
	@docker run \
		--rm \
		--name ${SERVICE_NAME} \
		--network host \
		${DOCKER_IMAGE}:${DOCKER_TAG}

.PHONY: test
test:
	@go test --cover -v ./...