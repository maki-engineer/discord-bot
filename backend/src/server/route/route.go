package route

import (
	_ "discord-bot/docs"
	"discord-bot/src/config"
	"discord-bot/src/presentation/member/handler"
	"discord-bot/src/presentation/middleware"
	"net/http"

	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRoutes(memberHandler *handler.MemberHandler, sessionService middleware.SessionService) *gin.Engine {
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

	// TODO: Discord認証処理を実装することになったら以下のAPIで実装していく
	// discord := r.Group("/discord")

	// {
	// 	discord.GET("/auth")
	// 	discord.GET("/auth/callback")
	// }

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
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", config.LoadConfig().AllowURL)
		c.Writer.Header().Set("Vary", "Origin")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	}
}
