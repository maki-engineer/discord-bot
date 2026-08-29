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

// GetMembersByBirthdayMonth godoc
// @Summary 指定された月が誕生日のメンバー一覧を取得するAPI
// @Tags members
// @Accept json
// @Produce json
// @Param birthday_month query int true "Birthday month (1-12)"
// @Success 200 {object} dto.GetMembersByBirthdayMonthResponse
// @Failure 400 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /members [get]
func (h *MemberHandler) GetMembersByBirthdayMonth(c *gin.Context) {
	var req memberDto.GetMembersByBirthdayMonthRequest

	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, errorDto.ErrorResponse{
			Result:  "error",
			Message: err.Error(),
		})
		return
	}

	ctx := c.Request.Context()

	members, err := h.useCase.GetMembersByBirthdayMonth(ctx, req.BirthdayMonth)
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
