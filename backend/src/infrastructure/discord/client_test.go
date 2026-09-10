package discord

import (
	"context"
	"io"
	"net/http"
	"strings"
	"testing"

	"discord-bot/src/config"
	"discord-bot/src/domain/auth"
)

type roundTripFunc func(*http.Request) (*http.Response, error)

func (f roundTripFunc) RoundTrip(request *http.Request) (*http.Response, error) {
	return f(request)
}

func TestClient_ExchangeCode(t *testing.T) {
	client := NewClient(config.Config{
		DiscordClientID:     "client-id",
		DiscordClientSecret: "client-secret",
		DiscordRedirectURI:  "http://localhost:8080/discord/auth/callback",
	})
	client.httpClient = &http.Client{Transport: roundTripFunc(func(request *http.Request) (*http.Response, error) {
		body, _ := io.ReadAll(request.Body)
		form := string(body)
		if request.Method != http.MethodPost || !strings.Contains(form, "code=code") {
			t.Fatalf("unexpected token request: %s %s", request.Method, form)
		}
		return jsonResponse(http.StatusOK, `{"access_token":"access-token"}`), nil
	})}

	token, err := client.ExchangeCode(context.Background(), "code")
	if err != nil {
		t.Fatalf("ExchangeCode returned an error: %v", err)
	}
	if token != "access-token" {
		t.Fatalf("expected access token, got %q", token)
	}
}

func TestClient_VerifyGuildMemberReturnsNotMember(t *testing.T) {
	client := NewClient(config.Config{
		DiscordGuildID:  "guild-id",
		DiscordBotToken: "bot-token",
	})
	client.httpClient = &http.Client{Transport: roundTripFunc(func(request *http.Request) (*http.Response, error) {
		if request.Header.Get("Authorization") != "Bot bot-token" {
			t.Fatalf("expected bot authorization header, got %q", request.Header.Get("Authorization"))
		}
		return jsonResponse(http.StatusNotFound, `{}`), nil
	})}

	err := client.VerifyGuildMember(context.Background(), "user-id")
	if err != auth.ErrNotGuildMember {
		t.Fatalf("expected ErrNotGuildMember, got %v", err)
	}
}

func jsonResponse(status int, body string) *http.Response {
	return &http.Response{
		StatusCode: status,
		Body:       io.NopCloser(strings.NewReader(body)),
		Header:     make(http.Header),
	}
}
