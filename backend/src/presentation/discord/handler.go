package discord

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"discord-bot/src/domain/auth"
	"discord-bot/src/presentation/dto"

	"github.com/gin-gonic/gin"
)

type AuthUseCase interface {
	Login(ctx context.Context, code string) (string, error)
}

type OAuthConfig struct {
	ClientID    string
	RedirectURI string
	FrontendURL string
}

type Handler struct {
	config  OAuthConfig
	useCase AuthUseCase
}

func NewHandler(cfg OAuthConfig, useCase AuthUseCase) *Handler {
	return &Handler{config: cfg, useCase: useCase}
}

// Auth godoc
// @Summary Discord OAuth2 認証API
// @Tags discord
// @Produce json
// @Success 302
// @Failure 500 {object} dto.ErrorResponse
// @Router /discord/auth [get]
func (h *Handler) Auth(c *gin.Context) {
	state, err := randomToken(32)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Result:  "error",
			Message: err.Error(),
		})
		return
	}

	secure := strings.HasPrefix(h.config.RedirectURI, "https://")
	if secure {
		c.SetSameSite(http.SameSiteNoneMode)
	}
	c.SetCookie("discord_oauth_state", state, 600, "/", "", secure, true)
	query := url.Values{
		"client_id":     []string{h.config.ClientID},
		"redirect_uri":  []string{h.config.RedirectURI},
		"response_type": []string{"code"},
		"scope":         []string{"identify"},
		"state":         []string{state},
	}
	c.Redirect(http.StatusFound, "https://discord.com/oauth2/authorize?"+query.Encode())
}

// Callback godoc
// @Summary Discord OAuth2 認証コールバックAPI
// @Tags discord
// @Produce json
// @Param code query string false "Discord OAuth2認証コード"
// @Param state query string false "OAuth2 state"
// @Param error query string false "Discord OAuth2認証エラー"
// @Success 302
// @Failure 302
// @Router /discord/auth/callback [get]
func (h *Handler) Callback(c *gin.Context) {
	frontendURL := h.config.FrontendURL
	if frontendURL == "" {
		frontendURL = "http://localhost:3000"
	}
	if c.Query("error") != "" {
		redirectWithError(c, frontendURL, "discord_auth_denied")
		return
	}

	state, err := c.Cookie("discord_oauth_state")
	if err != nil || state == "" || state != c.Query("state") {
		redirectWithError(c, frontendURL, "invalid_oauth_state")
		return
	}
	c.SetCookie("discord_oauth_state", "", -1, "/", "", false, true)

	code := c.Query("code")
	if code == "" {
		redirectWithError(c, frontendURL, "missing_oauth_code")
		return
	}

	sessionID, err := h.useCase.Login(c.Request.Context(), code)
	if err != nil {
		if errors.Is(err, auth.ErrNotGuildMember) {
			redirectWithError(c, frontendURL, "not_a_guild_member")
			return
		}
		redirectWithError(c, frontendURL, "discord_login_failed")
		return
	}

	secure := strings.HasPrefix(h.config.RedirectURI, "https://")
	if secure {
		c.SetSameSite(http.SameSiteNoneMode)
	}
	c.SetCookie("session_id", sessionID, 86400, "/", "", secure, true)
	c.Redirect(http.StatusFound, frontendURL+"/")
}

func redirectWithError(c *gin.Context, frontendURL string, errorCode string) {
	c.Redirect(http.StatusFound, frontendURL+"/?error="+url.QueryEscape(errorCode))
}

func randomToken(size int) (string, error) {
	bytes := make([]byte, size)
	if _, err := rand.Read(bytes); err != nil {
		return "", fmt.Errorf("generate random token: %w", err)
	}
	return hex.EncodeToString(bytes), nil
}
