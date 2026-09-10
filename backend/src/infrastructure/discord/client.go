package discord

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"

	"discord-bot/src/config"
	"discord-bot/src/domain/auth"
)

type Client struct {
	config     config.Config
	httpClient *http.Client
}

type tokenResponse struct {
	AccessToken string `json:"access_token"`
}

type userResponse struct {
	ID string `json:"id"`
}

func NewClient(cfg config.Config) *Client {
	return &Client{config: cfg, httpClient: http.DefaultClient}
}

func (c *Client) ExchangeCode(ctx context.Context, code string) (string, error) {
	form := url.Values{
		"client_id":     []string{c.config.DiscordClientID},
		"client_secret": []string{c.config.DiscordClientSecret},
		"grant_type":    []string{"authorization_code"},
		"code":          []string{code},
		"redirect_uri":  []string{c.config.DiscordRedirectURI},
	}
	request, err := http.NewRequestWithContext(ctx, http.MethodPost, "https://discord.com/api/oauth2/token", strings.NewReader(form.Encode()))
	if err != nil {
		return "", err
	}
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	var response tokenResponse
	if err := c.doJSON(request, &response); err != nil {
		log.Printf("discord ExchangeCode: doJSON failed: %v", err)
		return "", err
	}
	log.Println("discord ExchangeCode: response success")
	if response.AccessToken == "" {
		return "", fmt.Errorf("discord returned an empty access token")
	}
	return response.AccessToken, nil
}

func (c *Client) GetUser(ctx context.Context, accessToken string) (*auth.DiscordUser, error) {
	request, err := c.newAuthenticatedRequest(ctx, http.MethodGet, "https://discord.com/api/users/@me", accessToken, "Bearer")
	if err != nil {
		return nil, err
	}

	var response userResponse
	if err := c.doJSON(request, &response); err != nil {
		return nil, err
	}
	return &auth.DiscordUser{ID: response.ID}, nil
}

func (c *Client) VerifyGuildMember(ctx context.Context, userID string) error {
	endpoint := "https://discord.com/api/guilds/" + url.PathEscape(c.config.DiscordGuildID) + "/members/" + url.PathEscape(userID)
	request, err := c.newAuthenticatedRequest(ctx, http.MethodGet, endpoint, c.config.DiscordBotToken, "Bot")
	if err != nil {
		return err
	}

	response, err := c.httpClient.Do(request)
	if err != nil {
		return err
	}

	defer func() {
		_ = response.Body.Close()
	}()

	if response.StatusCode == http.StatusNotFound {
		return auth.ErrNotGuildMember
	}
	if response.StatusCode < http.StatusOK || response.StatusCode >= http.StatusMultipleChoices {
		return fmt.Errorf("discord guild member lookup returned status %d", response.StatusCode)
	}
	return nil
}

func (c *Client) newAuthenticatedRequest(ctx context.Context, method string, endpoint string, token string, scheme string) (*http.Request, error) {
	request, err := http.NewRequestWithContext(ctx, method, endpoint, nil)
	if err != nil {
		return nil, err
	}
	request.Header.Set("Authorization", scheme+" "+token)
	return request, nil
}

func (c *Client) doJSON(request *http.Request, target interface{}) error {
	log.Println("discord doJSON: request start")
	response, err := c.httpClient.Do(request)
	if err != nil {
		log.Printf("discord doJSON: request failed: %v", err)
		return err
	}

	defer func() {
		_ = response.Body.Close()
	}()

	log.Printf("discord doJSON: response received, status=%d", response.StatusCode)

	if response.StatusCode < http.StatusOK || response.StatusCode >= http.StatusMultipleChoices {
		log.Printf("discord doJSON: unexpected status=%d", response.StatusCode)
		return fmt.Errorf("discord API returned status %d", response.StatusCode)
	}

	if err := json.NewDecoder(response.Body).Decode(target); err != nil {
		log.Printf("discord doJSON: JSON decode failed: %v", err)
		return err
	}

	log.Println("discord doJSON: JSON decode success")

	return nil
}
