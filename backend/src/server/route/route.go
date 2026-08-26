package route

import (
	"discord-bot/src/presentation/member/handler"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(memberHandler *handler.MemberHandler) *gin.Engine {
	r := gin.Default()

	api := r.Group("/api")
	members := api.Group("/members")

	{
		members.GET("", memberHandler.GetMembersByBirthdayMonth)
	}

	return r
}
