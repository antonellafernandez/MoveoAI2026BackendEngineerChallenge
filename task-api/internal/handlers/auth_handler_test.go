package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"

	"task-api/internal/config"
	"task-api/internal/dto"
)

func setupAuthRouter(cfg *config.Config) *gin.Engine {
	gin.SetMode(gin.TestMode)

	r := gin.New()
	r.POST("/login", func(c *gin.Context) {
		Login(c, cfg)
	})

	return r
}

// TestLogin_Success tests the successful login scenario
func TestLogin_Success(t *testing.T) {
	cfg := &config.Config{
		Auth: config.AuthConfig{
			AdminUsername: "admin",
			AdminPassword: "admin",
		},
		JWT: config.JWTConfig{
			Secret: "test-secret",
		},
	}

	r := setupAuthRouter(cfg)

	body := dto.LoginRequest{
		Username: "admin",
		Password: "admin",
	}

	jsonBody, _ := json.Marshal(body)

	req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected 200, got %d", w.Code)
	}

	var resp dto.LoginResponse
	err := json.Unmarshal(w.Body.Bytes(), &resp)
	if err != nil {
		t.Errorf("invalid response body")
	}

	if resp.Token == "" {
		t.Errorf("expected token, got empty")
	}
}

// TestLogin_InvalidCredentials tests the login scenario with invalid credentials
func TestLogin_InvalidCredentials(t *testing.T) {
	cfg := &config.Config{
		Auth: config.AuthConfig{
			AdminUsername: "admin",
			AdminPassword: "admin",
		},
		JWT: config.JWTConfig{
			Secret: "test-secret",
		},
	}

	r := setupAuthRouter(cfg)

	body := dto.LoginRequest{
		Username: "wrong",
		Password: "wrong",
	}

	jsonBody, _ := json.Marshal(body)

	req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	if w.Code != http.StatusUnauthorized {
		t.Errorf("expected 401, got %d", w.Code)
	}

	var resp dto.MessageResponse
	_ = json.Unmarshal(w.Body.Bytes(), &resp)

	if resp.Message != "invalid credentials" {
		t.Errorf("unexpected message: %s", resp.Message)
	}
}

// TestLogin_BadRequest tests the login scenario with an invalid request body
func TestLogin_BadRequest(t *testing.T) {
	cfg := &config.Config{
		Auth: config.AuthConfig{
			AdminUsername: "admin",
			AdminPassword: "admin",
		},
		JWT: config.JWTConfig{
			Secret: "test-secret",
		},
	}

	r := setupAuthRouter(cfg)

	req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer([]byte(`{}`)))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("expected 400, got %d", w.Code)
	}
}
