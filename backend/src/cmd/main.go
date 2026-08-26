package main

import (
	"discord-bot/src/application/member/usecase"
	"discord-bot/src/config"
	"discord-bot/src/infrastructure/db"
	"discord-bot/src/infrastructure/repository"
	"discord-bot/src/presentation/member/handler"
	"discord-bot/src/server/route"
)

func main() {
	config := config.LoadConfig()
	db, err := db.NewDB(config)
	if err != nil {
		panic(err)
	}

	repository := repository.NewMemberRepository(db)
	useCase := usecase.NewMemberUseCase(repository)
	handler := handler.NewMemberHandler(useCase)

	r := route.SetupRoutes(handler)

	r.Run(":8080")
}
