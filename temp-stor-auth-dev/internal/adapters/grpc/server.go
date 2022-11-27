package grpc

import (
	"context"
	"fmt"
	"net"

	"github.com/seggga/temp-stor-auth/internal/adapters/grpc/apigrpcsrv"
	"github.com/seggga/temp-stor-auth/internal/adapters/grpc/grpcauth"
	"github.com/seggga/temp-stor-auth/internal/domain/models"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Service struct {
	//stor ports.UserStorage
	//cfg.JWT.Secret, cfg.JWT.Duration, logger,cfg.GrpcPorg api-grpc-srv
	grpcServer *grpc.NewServer
	server     *apigrpcsrv.MyServer
	logger     *zap.Logger
	listener   net.Listener
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

	s.grpcServer = &grpc.NewServer()
	s.server = &apigrpcsrv.MyServer{
		secret:   secret,
		duration: duration,
	}

	grpcauth.RegisterMessageServiceServer(s.grpcServer, s.server)

	// register the reflection service which allows clients to determine the methods
	// for this gRPC service
	reflection.Register(s.grpcServer)

	return s
}

func (s *Service) Start(ctx context.Context) error {
	s.logger.Debug("starting GRPC server ...")

	err := s.grpcServer.Serve(s.listener)
	if err != nil {
		return fmt.Errorf("cannot start GRPC server: %v", err)
	}
	return nil
}

func (s *Service) Stop(ctx context.Context) error {
	s.logger.Debug("stopping REST server ...")
	return s.grpcServer.Shutdown(ctx)
}

func (s *Service) Validate(ctx context.Context, tokens *models.Token) error {

	return nil
}
