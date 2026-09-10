package discord

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"discord-bot/src/domain/auth"

	"github.com/gin-gonic/gin"
)

type mockAuthUseCase struct {
	sessionID string
	err       error
	code      string
}

func (m *mockAuthUseCase) Login(ctx context.Context, code string) (string, error) {
	m.code = code
	return m.sessionID, m.err
}

func TestHandler_AuthRedirectsToDiscord(t *testing.T) {
	gin.SetMode(gin.TestMode)
	handler := NewHandler(OAuthConfig{
		ClientID:    "client-id",
		RedirectURI: "http://localhost:8080/discord/auth/callback",
	}, &mockAuthUseCase{})
	router := gin.New()
	router.GET("/discord/auth", handler.Auth)

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, httptest.NewRequest(http.MethodGet, "/discord/auth", nil))

	if recorder.Code != http.StatusFound {
		t.Fatalf("expected status %d, got %d", http.StatusFound, recorder.Code)
	}
	location := recorder.Header().Get("Location")
	if !strings.HasPrefix(location, "https://discord.com/oauth2/authorize?") {
		t.Fatalf("unexpected redirect location: %s", location)
	}
	if !strings.Contains(location, "client_id=client-id") || !strings.Contains(location, "state=") {
		t.Fatalf("expected OAuth query parameters: %s", location)
	}
}

func TestHandler_CallbackCreatesSessionCookie(t *testing.T) {
	gin.SetMode(gin.TestMode)
	useCase := &mockAuthUseCase{sessionID: "session-id"}
	handler := NewHandler(OAuthConfig{
		RedirectURI: "http://localhost:8080/discord/auth/callback",
		FrontendURL: "http://localhost:3000",
	}, useCase)
	router := gin.New()
	router.GET("/discord/auth/callback", handler.Callback)

	recorder := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodGet, "/discord/auth/callback?state=expected&code=code", nil)
	request.AddCookie(&http.Cookie{Name: "discord_oauth_state", Value: "expected"})
	router.ServeHTTP(recorder, request)

	if recorder.Code != http.StatusFound {
		t.Fatalf("expected status %d, got %d", http.StatusFound, recorder.Code)
	}
	if useCase.code != "code" {
		t.Fatalf("expected authorization code to be passed, got %q", useCase.code)
	}
	if !containsCookie(recorder.Header().Values("Set-Cookie"), "session_id=session-id") {
		t.Fatalf("expected session cookie, got %v", recorder.Header().Values("Set-Cookie"))
	}
	if recorder.Header().Get("Location") != "http://localhost:3000/" {
		t.Fatalf("unexpected redirect: %s", recorder.Header().Get("Location"))
	}
}

func TestHandler_CallbackRejectsNonMember(t *testing.T) {
	gin.SetMode(gin.TestMode)
	handler := NewHandler(OAuthConfig{FrontendURL: "http://localhost:3000"}, &mockAuthUseCase{
		err: errors.Join(errors.New("membership check failed"), auth.ErrNotGuildMember),
	})
	router := gin.New()
	router.GET("/discord/auth/callback", handler.Callback)

	recorder := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodGet, "/discord/auth/callback?state=expected&code=code", nil)
	request.AddCookie(&http.Cookie{Name: "discord_oauth_state", Value: "expected"})
	router.ServeHTTP(recorder, request)

	if recorder.Code != http.StatusFound {
		t.Fatalf("expected status %d, got %d", http.StatusFound, recorder.Code)
	}
	if recorder.Header().Get("Location") != "http://localhost:3000/?error=not_a_guild_member" {
		t.Fatalf("unexpected redirect: %s", recorder.Header().Get("Location"))
	}
	if containsCookie(recorder.Header().Values("Set-Cookie"), "session_id=") {
		t.Fatal("expected no session cookie for a non-member")
	}
}

func containsCookie(cookies []string, value string) bool {
	for _, cookie := range cookies {
		if strings.Contains(cookie, value) {
			return true
		}
	}
	return false
}
