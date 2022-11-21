PROJECT?=github.com/LightAlykard/TempExch
PROJECTNAME=$(shell basename "$(PROJECT)")

TARGETOS?=linux
TARGETARCH?=amd64

CGO_ENABLED=0

RELEASE := $(shell git tag -l | tail -1 | grep -E "v.+"|| echo devel)
COMMIT := git-$(shell git rev-parse --short HEAD)
BUILD_TIME := $(shell date -u '+%Y-%m-%d_%H:%M:%S')
COPYRIGHT := "sanya-spb"

## build: Build test-env
build:
	GOOS=${TARGETOS} GOARCH=${TARGETARCH} CGO_ENABLED=${CGO_ENABLED} go build \
		-ldflags "-s -w \
		-X ${PROJECT}/pkg/version.version=${RELEASE} \
		-X ${PROJECT}/pkg/version.commit=${COMMIT} \
		-X ${PROJECT}/pkg/version.buildTime=${BUILD_TIME} \
		-X ${PROJECT}/pkg/version.copyright=${COPYRIGHT}" \
		-o ./bin/test-env ./cmd/test-env/

## image: Build test-env docker image
image:
	docker build -t otin-backend \
	--build-arg RELEASE=${RELEASE} \
	--build-arg COMMIT=${COMMIT} \
	--build-arg BUILD_TIME=${BUILD_TIME} \
	.
	@echo "\n\nTo start container:"
	@echo 'docker run -dit --restart unless-stopped -p 8080:8080 -v $(pwd)/conf:/app/data/conf --name test-env test-env:latest'

## run: Run test-env
run:
	go run ./cmd/test-env/ -config ./data/conf/config.yaml

## clean: Clean build files
clean:
	go clean
	rm -v ./bin/* 2> /dev/null || true

## test: Run unit test
test:
	go test -v -short ${PROJECT}/...

## integration: Run integration test
integration:
	 go test -v -run Integration ${PROJECT}/cmd/test-env/

## help: Show this
help: Makefile
	@echo " Choose a command run in "$(PROJECTNAME)":"
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'
