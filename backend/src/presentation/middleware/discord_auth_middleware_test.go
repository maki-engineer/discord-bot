package middleware

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"discord-bot/src/domain/auth"

	"github.com/gin-gonic/gin"
)

type mockSessionService struct {
	session *auth.Session
	err     error
}

func (m *mockSessionService) GetSessionByID(ctx context.Context, sessionID string) (*auth.Session, error) {
	return m.session, m.err
}

func (m *mockSessionService) CreateSession(ctx context.Context, session *auth.Session) error {
	return nil
}

func TestAuthMiddleware_AllowsValidSession(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.Use(AuthMiddleware(&mockSessionService{
		session: &auth.Session{SessionID: "session-id", ExpiresAt: time.Now().Add(time.Hour)},
	}))
	router.GET("/protected", func(c *gin.Context) { c.Status(http.StatusOK) })

	recorder := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodGet, "/protected", nil)
	request.AddCookie(&http.Cookie{Name: "session_id", Value: "session-id"})
	router.ServeHTTP(recorder, request)

	if recorder.Code != http.StatusOK {
		t.Fatalf("expected status %d, got %d", http.StatusOK, recorder.Code)
	}
}

func TestAuthMiddlewareRejectsMissingSession(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.Use(AuthMiddleware(&mockSessionService{}))
	router.GET("/protected", func(c *gin.Context) { c.Status(http.StatusOK) })

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, httptest.NewRequest(http.MethodGet, "/protected", nil))

	if recorder.Code != http.StatusUnauthorized {
		t.Fatalf("expected status %d, got %d", http.StatusUnauthorized, recorder.Code)
	}
}

func TestAuthMiddlewareRejectsInvalidSession(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.Use(AuthMiddleware(&mockSessionService{err: errors.New("session not found")}))
	router.GET("/protected", func(c *gin.Context) { c.Status(http.StatusOK) })

	recorder := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodGet, "/protected", nil)
	request.AddCookie(&http.Cookie{Name: "session_id", Value: "invalid"})
	router.ServeHTTP(recorder, request)

	if recorder.Code != http.StatusUnauthorized {
		t.Fatalf("expected status %d, got %d", http.StatusUnauthorized, recorder.Code)
	}
}
