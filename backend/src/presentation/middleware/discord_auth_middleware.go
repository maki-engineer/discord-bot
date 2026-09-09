package middleware

import (
	errorDto "discord-bot/src/presentation/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(sessionService SessionService) gin.HandlerFunc {
	return func(c *gin.Context) {
		sessionID, err := c.Cookie("session_id")
		if err != nil || sessionID == "" {
			c.JSON(http.StatusUnauthorized, errorDto.ErrorResponse{
				Result:  "error",
				Message: err.Error(),
			})
			c.Abort()
			return
		}

		ctx := c.Request.Context()
		session, err := sessionService.GetSessionByID(ctx, sessionID)
		if err != nil || session == nil {
			c.JSON(http.StatusUnauthorized, errorDto.ErrorResponse{
				Result:  "error",
				Message: err.Error(),
			})
			c.Abort()
			return
		}

		c.Next()
	}
}
