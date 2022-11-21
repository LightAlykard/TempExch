package memory

import (
	"context"
	"errors"
	"sync"

	"github.com/google/uuid"
	"github.com/seggga/temp-stor-auth/internal/domain/models"
)

var (
	NOT_FOUND = errors.New("user with specified login was not found")
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
		return nil, NOT_FOUND
	}
	return &u, nil
}
