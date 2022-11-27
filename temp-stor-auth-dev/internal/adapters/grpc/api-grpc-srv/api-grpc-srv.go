package apigrpcsrv

import (
	"context"
	"fmt"

	"github.com/dgrijalva/jwt-go/v4"

	"github.com/seggga/temp-stor-auth/internal/adapters/grpc/grpcauth"
)

var _ grpcauth.MessageServiceServer = &MyServer{}

type MyServer struct {
	//mainCtx context.Context
	secret   string
	duration int
	grpcauth.UnimplementedMessageServiceServer
}

func (s *MyServer) ServeJWT(ctx context.Context, msg *grpcauth.Message) (*grpcauth.Reply, error) {
	token, err := jwt.Parse(msg.body, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return s.secret, nil
	})

	if token.Valid {
		return &grpcauth.ReplyJWT{status: "OK"}, nil
	}

	return nil, err
}
