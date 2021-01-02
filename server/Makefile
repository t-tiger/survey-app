NAME    := survey-server
SRCS    := $(shell find . -type f -name '*.go')
LDFLAGS := -ldflags="-s -w -extldflags \"-static\""

bin/$(NAME): $(SRCS)
	@go build -tags netgo -installsuffix netgo $(LDFLAGS) -o bin/$(NAME)

.PHONY: install
install:
	go install $(LDFLAGS)

.PHONY: mod-download
mod-download:
	go mod download

.PHONY: go-vet
go-vet:
	go vet ./...

.PHONY: go-test
test:
	go test ./... -v

.PHONY: docker-base-build
docker-base-build:
	DOCKER_BUILDKIT=1 docker build \
		--target=base \
		-t survey-server-base .

.PHONY: docker-build
docker-build:
	DOCKER_BUILDKIT=1 docker build \
		--target=app \
		-t survey-server .
