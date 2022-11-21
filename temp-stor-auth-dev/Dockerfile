# downdload modules
FROM golang:1.17-buster as builder

WORKDIR /app
COPY go.* ./
RUN go mod download
COPY . ./
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -o auth-srv ./cmd/server/main.go

# copy binary+configs and start application
FROM alpine:3.15.4

RUN mkdir -p /app/configs
WORKDIR /app
COPY --from=builder /app/auth-srv ./
COPY --from=builder /app/configs/. ./configs/

CMD ["/app/auth-srv"]