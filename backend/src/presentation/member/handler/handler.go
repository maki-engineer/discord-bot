package handler

import (
	errorDto "discord-bot/src/presentation/dto"
	memberDto "discord-bot/src/presentation/member/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)

type MemberHandler struct {
	useCase MemberUseCase
}

func NewMemberHandler(useCase MemberUseCase) *MemberHandler {
	return &MemberHandler{
		useCase: useCase,
	}
}

func (h *MemberHandler) GetMembersByBirthdayMonth(c *gin.Context) {
	var req memberDto.GetMembersByBirthdayMonthRequest

	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, errorDto.ErrorResponse{
			Result:  "error",
			Message: err.Error(),
		})
		return
	}

	members, err := h.useCase.GetMembersByBirthdayMonth(req.BirthdayMonth)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorDto.ErrorResponse{
			Result:  "error",
			Message: err.Error(),
		})
		return
	}

	membersResponse := make([]memberDto.MemberBirthday, 0, len(members))

	for _, member := range members {
		membersResponse = append(membersResponse, memberDto.MemberBirthday{
			Name:  member.Name,
			Month: member.Month,
			Date:  member.Date,
		})
	}

	c.JSON(http.StatusOK, memberDto.GetMembersByBirthdayMonthResponse{
		Result:  "success",
		Members: membersResponse,
	})
}
