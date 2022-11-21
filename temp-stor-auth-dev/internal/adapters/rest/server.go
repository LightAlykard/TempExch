package rest

import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/seggga/temp-stor-auth/internal/domain/models"
	"github.com/seggga/temp-stor-auth/internal/ports"
	"go.uber.org/zap"
)

type Service struct {
	auth     ports.Auther
	server   *http.Server
	logger   *zap.Logger
	listener net.Listener

	// stor ports.UserStorage
}

func New(auth ports.Auther, logger *zap.Logger, port string) *Service {
	var err error
	s := &Service{
		auth:   auth,
		logger: logger,
	}

	s.listener, err = net.Listen("tcp", ":"+port)
	if err != nil {
		logger.Sugar().Fatalf("error creating listener on port %s: %v", port, err)
	}
	s.server = &http.Server{
		Handler: s.routes(),
	}

	return s
}

func (s *Service) Validate(ctx context.Context, token *models.Token) error {
	return nil
}

func (s *Service) Login(ctx context.Context, user, password string) (*models.Token, error) {
	return nil, nil
}

func (s *Service) Logout(ctx context.Context, token models.Token) {}

// Start starts REST service
func (s *Service) Start(ctx context.Context) error {
	s.logger.Debug("starting REST server ...")

	if err := s.server.Serve(s.listener); !errors.Is(err, http.ErrServerClosed) {
		return fmt.Errorf("cannot start REST server: %v", err)
	}
	return nil
}

// Stop grecefuly terminages the REST service
func (s *Service) Stop(ctx context.Context) error {
	s.logger.Debug("stopping REST server ...")
	return s.server.Shutdown(ctx)
}

func (s *Service) routes() http.Handler {
	r := chi.NewMux()

	r.Use(middleware.Heartbeat("/healthz"))
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))

	r.Mount("/", s.Handlers())

	return r
}
