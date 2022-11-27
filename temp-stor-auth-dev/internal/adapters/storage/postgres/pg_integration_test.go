package postgres

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/TempExch/temp-stor-auth-dev/internal/domain/models"
	"github.com/stretchr/testify/assert"
)

const (
	DSN = "postgres://root:pass@127.0.0.1:5432/test_db"
)

var store *Storage

type testCase struct {
	name      string
	login     string
	expectErr bool
	id        string
	hash      string
}

// test cases:
var testCases = []testCase{
	{
		name:      "1. proper login, no errors",
		login:     "first-user",
		expectErr: false,
		id:        "8b545c16-5a5f-4fe3-b721-a61cab06dca9",
		hash:      "first-hash",
	},
	{
		name:      "2. empty return (no user found)",
		login:     "fake-user",
		expectErr: true,
	},
}

func TestMain(m *testing.M) {
	store = New(DSN)
	os.Exit(m.Run())
}

func TestGet(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			u, err := store.Get(ctx, tc.login)
			if tc.expectErr {
				assert.Nil(t, u)
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
				assert.NotNil(t, u)
				assert.Equal(t, tc.id, u.ID.String())
				assert.Equal(t, tc.login, u.Name)
				assert.Equal(t, tc.hash, u.Hash)
			}
		})
	}
}

func TestDelete(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := store.Delete(ctx, tc.login)
			if tc.expectErr {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)

				_, err = store.Get(ctx, tc.login)
				assert.NotNil(t, err)
			}
		})
	}
}

func TestInsert(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	name, hash := "new-user", "new-hash"
	_, err := store.Insert(ctx, &models.User{
		Name: name,
		Hash: hash,
	})
	assert.Nil(t, err)

	u, err := store.Get(ctx, name)
	assert.Nil(t, err)
	assert.Equal(t, u.Name, name)
	assert.Equal(t, u.Hash, hash)
	assert.Nil(t, err)

	_, err = store.Delete(ctx, name)

	_, err = store.Insert(ctx, &models.User{
		Name: "",
		Hash: "",
	})
	assert.NotNil(t, err)

}
