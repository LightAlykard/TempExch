package grpc

import (
	"context"
	"fmt"
	"net"

	grpcServer "github.com/TempExch/temp-stor-auth-dev/internal/adapters/grpc/grpcServer"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type Service struct {
	newgrpcServer *grpc.Server
	server        *MygRPCServer
	logger        *zap.Logger
	listener      net.Listener
}

func New(secret string, duration int, logger *zap.Logger, port string) *Service {
	s := &Service{
		logger: logger,
	}

	var err error
	s.listener, err = net.Listen("tcp", ":"+port)
	if err != nil {
		logger.Sugar().Fatalf("failed to listen: %s: %v", port, err)
	}

	s.server = &MygRPCServer{
		secret:   secret,
		duration: duration,
	}

	// register the reflection service which allows clients to determine the methods
	// for this gRPC service
	//reflection.Register(newgrpcServer)

	return s
}

func (s *Service) Start(ctx context.Context) error {

	s.newgrpcServer = grpc.NewServer()
	grpcServer.RegisterJWTChekServiceServer(s.newgrpcServer, s.server)
	s.logger.Debug("starting GRPC server ...")

	err := s.newgrpcServer.Serve(s.listener)
	if err != nil {
		return fmt.Errorf("cannot start GRPC server: %v", err)
	}
	return nil
}

func (s *Service) Stop(ctx context.Context) {
	s.logger.Debug("stopping REST server ...")
	//return s.newgrpcServer.Shutdown(ctx)
	s.newgrpcServer.GracefulStop()
}
