#build stage
FROM golang:alpine AS builder
RUN apk add --no-cache git
WORKDIR /go/src/app
COPY . .
#RUN go mod download -x

ARG PROJECT=github.com/LightAlykard/TempExch
ARG TARGETOS=linux
ARG TARGETARCH=amd64
ARG RELEASE='0.0.0'
ARG COMMIT='2020-11-21'
ARG BUILD_TIME='23:30'
ARG COPYRIGHT="sanya-spb"
ARG CGO_ENABLED=0

RUN GOOS=${TARGETOS} GOARCH=${TARGETARCH} CGO_ENABLED=${CGO_ENABLED} go build \
    -ldflags "-s -w \
    -X ${PROJECT}/pkg/version.version=${RELEASE} \
    -X ${PROJECT}/pkg/version.commit=${COMMIT} \
    -X ${PROJECT}/pkg/version.buildTime=${BUILD_TIME} \
    -X ${PROJECT}/pkg/version.copyright=${COPYRIGHT}" \
    -o ./bin/test-env ./cmd/test-env/

#final stage
FROM alpine:latest
RUN apk --no-cache add ca-certificates
COPY --from=builder /go/src/app/bin/test-env /app/test-env
#COPY --from=builder /go/src/app/data /app/data
RUN adduser -SDH goapp
USER goapp
WORKDIR /app
ENTRYPOINT /app/test-env -config /app/data/conf/config.yaml -debug
LABEL Name=test-env
#VOLUME ["/app/data/conf", "/app/data/logs", "/app/data/ui"]
#EXPOSE 8080/tcp
