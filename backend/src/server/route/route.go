package route

import (
	_ "discord-bot/docs"
	"discord-bot/src/config"
	discordHandler "discord-bot/src/presentation/discord"
	"discord-bot/src/presentation/member/handler"
	"discord-bot/src/presentation/middleware"
	"net/http"

	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRoutes(memberHandler *handler.MemberHandler, sessionService middleware.SessionService, discordAuthHandler *discordHandler.Handler) *gin.Engine {
	r := gin.Default()
	r.Use(corsMiddleware())

	swagger := r.Group("/swagger", setupSwaggerAuth())

	{
		swagger.GET("/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	api := r.Group("/api")
	api.Use(middleware.AuthMiddleware(sessionService))
	members := api.Group("/members")

	{
		members.GET("", memberHandler.GetMembersByBirthdayMonth)
	}

	discord := r.Group("/discord")

	{
		discord.GET("/auth", discordAuthHandler.Auth)
		discord.GET("/auth/callback", discordAuthHandler.Callback)
	}

	return r
}

func setupSwaggerAuth() gin.HandlerFunc {
	username := config.LoadConfig().BasicUserName
	password := config.LoadConfig().BasicPassword

	swaggerAuth := gin.BasicAuth(gin.Accounts{
		username: password,
	})

	return swaggerAuth
}

func corsMiddleware() gin.HandlerFunc {
	allowURL := config.LoadConfig().AllowURL

	if allowURL == "" {
		allowURL = "http://localhost:3000"
	}

	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", allowURL)
		c.Writer.Header().Set("Vary", "Origin")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	}
}
