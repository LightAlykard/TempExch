package grpc

import (
	"context"
	"log"
	"net"
	"testing"
	"time"

	grpcServer "github.com/TempExch/temp-stor-auth-dev/internal/adapters/grpc/grpcServer"
	crT "github.com/TempExch/temp-stor-auth-dev/internal/domain/auth"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/grpc/test/bufconn"
)

func dialer() func(context.Context, string) (net.Conn, error) {
	listener := bufconn.Listen(1024 * 1024)

	server := grpc.NewServer()

	grpcServer.RegisterJWTChekServiceServer(server, &MygRPCServer{
		secret:   "super-secret-key",
		duration: 10,
	})

	go func() {
		if err := server.Serve(listener); err != nil {
			log.Fatal(err)
		}
	}()

	return func(context.Context, string) (net.Conn, error) {
		return listener.Dial()
	}
}

func TestJWTChekService_JWTchek(t *testing.T) {

	tokenInvalid, _ := crT.CreateToken("some-user", "super-secret-key", 1)
	time.Sleep(time.Second)
	tokenValid, _ := crT.CreateToken("some-user", "super-secret-key", 500)

	tests := []struct {
		name     string
		jwttoken string
		res      *grpcServer.JWTChekResponse
		errCode  codes.Code
		errMsg   string
	}{
		{
			"first request",
			tokenInvalid,
			&grpcServer.JWTChekResponse{Ok: false},
			codes.DeadlineExceeded,
			"Token is either expired or not active yet",
		},
		{
			"second request",
			tokenValid,
			&grpcServer.JWTChekResponse{Ok: true},
			codes.OK,
			"",
		},
	}

	ctx := context.Background()

	conn, err := grpc.DialContext(ctx, "", grpc.WithInsecure(), grpc.WithContextDialer(dialer()))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := grpcServer.NewJWTChekServiceClient(conn)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			request := &grpcServer.JWTChekRequest{Jwttoken: tt.jwttoken}

			response, err := client.JWTchek(ctx, request)

			if response != nil {
				if response.GetOk() != tt.res.GetOk() {
					t.Error("response: expected", tt.res.GetOk(), "received", response.GetOk())
				}
			}

			if err != nil {
				if er, ok := status.FromError(err); ok {
					if er.Code() != tt.errCode {
						t.Error("error code: expected", codes.InvalidArgument, "received", er.Code())
					}
					if er.Message() != tt.errMsg {
						t.Error("error message: expected", tt.errMsg, "received", er.Message())
					}
				}
			}
		})
	}
}
