package memory

import (
	"context"
	"sync"

	er "github.com/TempExch/temp-stor-auth-dev/internal/adapters/storage"
	"github.com/TempExch/temp-stor-auth-dev/internal/domain/models"
	"github.com/google/uuid"
)

type Storage struct {
	m     map[string]models.User
	mutex sync.Mutex
}

// New creates a new memory storage
func New() *Storage {

	m := make(map[string]models.User, 2)
	m["user1"] = models.User{
		ID:   uuid.New(),
		Name: "user1",
		Hash: "123",
	}
	m["user2"] = models.User{
		ID:   uuid.New(),
		Name: "user2",
		Hash: "456",
	}
	return &Storage{
		m:     m,
		mutex: sync.Mutex{},
	}
}

// Get looks for a user by login
func (s *Storage) Get(ctx context.Context, login string) (*models.User, error) {
	s.mutex.Lock()
	u, ok := s.m[login]
	s.mutex.Unlock()

	if !ok {
		return nil, er.NOT_FOUND
	}
	return &u, nil
}
