//go:generate protoc --go_out=./grpcauth --go_opt=paths=source_relative --go-grpc_out=./grpcauth --go-grpc_opt=paths=source_relative ./grpcauth.proto

package grpc
