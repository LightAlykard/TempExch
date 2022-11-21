package ports

import (
	"context"

	"github.com/seggga/temp-stor-auth/internal/domain/models"
)

type UserStorage interface {
	Get(ctx context.Context, login string) (*models.User, error)
}
