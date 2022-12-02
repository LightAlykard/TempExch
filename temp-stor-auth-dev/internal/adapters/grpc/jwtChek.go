package grpc

import (
	"context"
	"errors"

	"github.com/golang-jwt/jwt/v4"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/TempExch/temp-stor-auth-dev/internal/adapters/grpc/grpcServer"
)

var _ grpcServer.JWTChekServiceServer = &MygRPCServer{}

type MygRPCServer struct {
	//mainCtx context.Context
	secret   string
	duration int
	grpcServer.UnimplementedJWTChekServiceServer
}

func (s *MygRPCServer) JWTchek(ctx context.Context, req *grpcServer.JWTChekRequest) (*grpcServer.JWTChekResponse, error) {
	token, err := jwt.Parse(req.Jwttoken, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.secret), nil
	})

	if token.Valid {
		return &grpcServer.JWTChekResponse{Ok: true}, status.Error(codes.OK, "Token is valid")
	} else if errors.Is(err, jwt.ErrTokenMalformed) {
		return &grpcServer.JWTChekResponse{Ok: false}, status.Error(codes.InvalidArgument, "That's not even a token")
	} else if errors.Is(err, jwt.ErrTokenExpired) || errors.Is(err, jwt.ErrTokenNotValidYet) {
		return &grpcServer.JWTChekResponse{Ok: false}, status.Error(codes.DeadlineExceeded, "Token is either expired or not active yet")
	} else {
		return &grpcServer.JWTChekResponse{Ok: false}, status.Error(codes.InvalidArgument, "Couldn't handle this token")
	}
}
