package main

import (
	"discord-bot/src/application/auth/service"
	authUsecase "discord-bot/src/application/auth/usecase"
	"discord-bot/src/application/member/usecase"
	"discord-bot/src/config"
	"discord-bot/src/infrastructure/db"
	discordInfrastructure "discord-bot/src/infrastructure/discord"
	"discord-bot/src/infrastructure/repository"
	discordHandler "discord-bot/src/presentation/discord"
	"discord-bot/src/presentation/member/handler"
	"discord-bot/src/server/route"
	"log"
)

// @title 235Bot API
// @version 1.0
// @license.name Maki
// @description 235botのデータを取得・操作したり、認証処理をするためのAPI
// @BasePath /api
func main() {
	config := config.LoadConfig()
	db, err := db.NewDB(config)
	if err != nil {
		panic(err)
	}

	memberRepository := repository.NewMemberRepository(db)
	sessionRepository := repository.NewSessionRepository(db)
	useCase := usecase.NewMemberUseCase(memberRepository)
	handler := handler.NewMemberHandler(useCase)
	middleware := service.NewSessionService(sessionRepository)
	discordGateway := discordInfrastructure.NewClient(config)
	discordLoginUseCase := authUsecase.NewDiscordAuthUseCase(discordGateway, sessionRepository)
	discordAuthHandler := discordHandler.NewHandler(discordHandler.OAuthConfig{
		ClientID:    config.DiscordClientID,
		RedirectURI: config.DiscordRedirectURI,
		FrontendURL: config.FrontendURL,
	}, discordLoginUseCase)

	r := route.SetupRoutes(handler, middleware, discordAuthHandler)

	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
