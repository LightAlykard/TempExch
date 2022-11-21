package rest

import (
	"bytes"
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/seggga/temp-stor-auth/internal/domain/models"
	"go.uber.org/zap"
)

type loginCases struct {
	name       string
	body       string
	statusCode int
}

func TestLogin(t *testing.T) {
	auth := &mockAuther{}
	logger, err := zap.NewDevelopment()
	if err != nil {
		t.Fatal("cannot create test logger")
	}
	s := New(auth, logger, "")

	testCases := []loginCases{
		{
			name:       "1. empyt body",
			body:       "",
			statusCode: 400,
		},
		{
			name:       "2. incorrect body format",
			body:       "{}",
			statusCode: 401,
		},
		{
			name:       "3. incorrect body format",
			body:       `{"field1":"123","field2":"123"}`,
			statusCode: 401,
		},
		{
			name:       "4. incorrect login/pass",
			body:       `{"username":"bad-login","password":"bad-pass"}`,
			statusCode: 401,
		},
		{
			name:       "5. good login/pass",
			body:       `{"username":"good-login","password":"good-pass"}`,
			statusCode: 200,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			body := bytes.NewReader([]byte(tc.body))
			req, err := http.NewRequest("POST", "/login", body)
			if err != nil {
				t.Errorf("Error creating a new request: %v", err)
			}

			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(s.login)
			handler.ServeHTTP(rr, req)

			if status := rr.Code; status != tc.statusCode {
				t.Errorf("wrong status code, expected: %d, got: %d.", tc.statusCode, status)
			}

		})
	}

}

type mockAuther struct{}

func (t *mockAuther) Validate(ctx context.Context, token models.Token) (string, error) {
	select {
	case <-ctx.Done():
		return "", errors.New("context closed")
	default:
		if token.Access == "good-token" {
			return "good", nil
		}
		return "", errors.New("bad token")
	}
}

func (t *mockAuther) Login(ctx context.Context, login, password string) (*models.Token, error) {
	select {
	case <-ctx.Done():
		return nil, errors.New("context closed")
	default:
		if login == "good-login" && password == "good-pass" {
			return &models.Token{
				Access: "good-token",
			}, nil
		}
		return nil, errors.New("bad login")
	}
}
