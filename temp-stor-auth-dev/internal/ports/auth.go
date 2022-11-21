package ports

import (
	"context"

	"github.com/seggga/temp-stor-auth/internal/domain/models"
)

type Auther interface {
	Validate(ctx context.Context, token models.Token) (string, error)
	Login(ctx context.Context, login, password string) (*models.Token, error)
}
