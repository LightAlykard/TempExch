package postgres

import (
	"context"
	"errors"
	"fmt"
	"log"

	er "github.com/TempExch/temp-stor-auth-dev/internal/adapters/storage"
	"github.com/TempExch/temp-stor-auth-dev/internal/domain/models"
	"github.com/google/uuid"
	"github.com/jackc/pgx"
	"github.com/jackc/pgx/v4/pgxpool"
)

// Storage ...
type Storage struct {
	Pool *pgxpool.Pool
}

// New ...
func New(dsn string) (s *Storage) {
	ctx := context.Background()
	pool, err := pgxpool.Connect(ctx, dsn)
	if err != nil {
		log.Fatalf("unable to connect to database: %v", err)
	}

	if err = pool.Ping(ctx); err != nil {
		log.Fatalf("unable to ping database: %v", err)
	}
	s = &Storage{Pool: pool}
	return s
}

// Insert adds event about task that has not been stored yet
func (s *Storage) Insert(ctx context.Context, usr *models.User) (string, error) {

	if usr.Name == "" {
		return "", er.EmptyLogin
	}
	if usr.Hash == "" {
		return "", er.EmpthPass
	}
	_, err := s.Get(ctx, usr.Name)
	if err == nil {
		return "", er.UserExists
	}

	query := "INSERT INTO auth.users (id, name, hash) values ($1, $2, $3) RETURNING id"
	row := s.Pool.QueryRow(ctx, query,
		uuid.New(),
		usr.Name,
		usr.Hash,
	)
	var id uuid.UUID
	if err := row.Scan(&id); err != nil {
		return "", fmt.Errorf("error inserting new user in db: %v", err)
	}

	return id.String(), nil
}

// Get looks for a user by login
func (s *Storage) Get(ctx context.Context, login string) (*models.User, error) {
	query := "SELECT id, name, hash FROM auth.users WHERE name = $1"
	row := s.Pool.QueryRow(ctx, query, login)

	var usr models.User
	if err := row.Scan(&usr.ID, &usr.Name, &usr.Hash); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, er.NotFound
		}
		return nil, fmt.Errorf("error reading user from db: %v", err)
	}
	return &usr, nil
}

// Insert adds event about task that has not been stored yet
func (s *Storage) Delete(ctx context.Context, login string) (string, error) {

	if login == "" {
		return "", er.EmptyLogin
	}
	_, err := s.Get(ctx, login)
	if err != nil && !errors.Is(err, pgx.ErrNoRows) {
		return "", er.NoSuchUser
	}

	query := "DELETE FROM auth.users WHERE name = $1 RETURNING id"
	row := s.Pool.QueryRow(ctx, query, login)
	var id uuid.UUID
	if err := row.Scan(&id); err != nil {
		return "", fmt.Errorf("error deleting the user %s from db: %v", login, err)
	}
	return id.String(), nil
}
