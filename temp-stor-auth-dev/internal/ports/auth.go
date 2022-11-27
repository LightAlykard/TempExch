package ports

import (
	"context"

	"github.com/TempExch/temp-stor-auth-dev/internal/domain/models"
)

type Auther interface {
	Validate(ctx context.Context, token models.Token) (string, error)
	Login(ctx context.Context, login, password string) (*models.Token, error)
}
