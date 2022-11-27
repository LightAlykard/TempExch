package grpc

import (
	"context"
	"fmt"
	"time"

	"github.com/seggga/temp-stor-auth/internal/adapters/grpc/grpcauth"
	"github.com/seggga/temp-stor-auth/internal/domain/models"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type Client struct {
	opts   []grpc.DialOption
	addres string
	logger *zap.Logger
	//conn grpc.Dial
}

func NewClient(addres string, logger *zap.Logger) *Client {
	opts := []grpc.DialOption{
		grpc.WithInsecure(),
	}

	c := &Client{
		opts: opts,
	}

	c.addres = addres //FIXME: Need IP from configfile
	return c

}

func (c *Client) StartClient(ctx context.Context) (*grpcauth.NewMessageServiceClient, error) {
	c.logger.Debug("starting REST server ...")
	var err error
	conn, err := grpc.Dial("127.0.0.1:5300", c.opts...)
	if err != nil {
		return fmt.Errorf("cannot start GRPC client: %v", err)
	}
	defer conn.Close()

	client := grpcauth.NewMessageServiceClient(conn)
	return client
}

func (c *Client) SendGRPC(client *grpcauth.NewMessageServiceClient, user models.User) (string, error) {
	request := &grpcauth.MessageAnal{
		ID:   user.ID,
		Name: user.Name,
		Hash: user.Hash,
		Time: time.Now(),
	}
	response, err := client.Do(context.Background(), request)
	if err != nil {
		c.logger.Debug("GRPC not delivered ...")
		return "NOT OK", err
	}

	return "OK", nil
}
