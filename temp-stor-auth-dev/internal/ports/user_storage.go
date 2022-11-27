package ports

import (
	"context"

	"github.com/TempExch/temp-stor-auth-dev/internal/domain/models"
)

type UserStorage interface {
	Get(ctx context.Context, login string) (*models.User, error)
}
