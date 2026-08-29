package route

import (
	_ "discord-bot/docs"
	"discord-bot/src/presentation/member/handler"

	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRoutes(memberHandler *handler.MemberHandler) *gin.Engine {
	r := gin.Default()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api := r.Group("/api")
	members := api.Group("/members")

	{
		members.GET("", memberHandler.GetMembersByBirthdayMonth)
	}

	return r
}
