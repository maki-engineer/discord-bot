package handler

import (
	"discord-bot/src/application/member/usecase"
	"discord-bot/src/presentation/member/dto"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/gin-gonic/gin"
)

type MockMemberUseCase struct {
	members []usecase.MemberBirthdayOutputData
	err     error
}

func (m *MockMemberUseCase) GetMembersByBirthdayMonth(birthdayMonth int) ([]usecase.MemberBirthdayOutputData, error) {
	return m.members, m.err
}

func TestMemberHandler_GetMembersByBirthdayMonth(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockUseCase := &MockMemberUseCase{
		members: []usecase.MemberBirthdayOutputData{
			{Name: "Alice", Month: 5, Date: 15},
			{Name: "Bob", Month: 5, Date: 20},
		},
		err: nil,
	}

	handler := NewMemberHandler(mockUseCase)

	router := gin.New()
	router.GET("/api/member", handler.GetMembersByBirthdayMonth)

	req := httptest.NewRequest(http.MethodGet, "/api/member?birthday_month=5", nil)

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, req)

	if recorder.Code != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, recorder.Code)
	}

	var response dto.GetMembersByBirthdayMonthResponse
	if err := json.Unmarshal(
		recorder.Body.Bytes(),
		&response,
	); err != nil {
		t.Errorf("Failed to unmarshal response: %v", err)
	}

	if response.Result != "success" {
		t.Errorf("Expected result 'success', but got '%s'", response.Result)
	}

	if len(response.Members) != 2 {
		t.Errorf("Expected 2 members, but got %d", len(response.Members))
	}

	expectedMembers := []dto.MemberBirthday{
		{Name: "Alice", Month: 5, Date: 15},
		{Name: "Bob", Month: 5, Date: 20},
	}

	if !reflect.DeepEqual(response.Members, expectedMembers) {
		t.Errorf("Expected members %v, but got %v", expectedMembers, response.Members)
	}
}

func TestMemberHandler_GetMembersByBirthdayMonth_InvalidQueryParam(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockUseCase := &MockMemberUseCase{}

	handler := NewMemberHandler(mockUseCase)

	router := gin.New()
	router.GET("/api/member", handler.GetMembersByBirthdayMonth)

	req := httptest.NewRequest(http.MethodGet, "/api/member?birthday_month=five", nil)

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, req)

	if recorder.Code != http.StatusBadRequest {
		t.Errorf("Expected status code %d, but got %d", http.StatusBadRequest, recorder.Code)
	}
}

func TestMemberHandler_GetMembersByBirthdayMonth_UseCaseError(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockUseCase := &MockMemberUseCase{
		err: errors.New("usecase error"),
	}

	handler := NewMemberHandler(mockUseCase)

	router := gin.New()
	router.GET("/api/member", handler.GetMembersByBirthdayMonth)

	req := httptest.NewRequest(http.MethodGet, "/api/member?birthday_month=5", nil)

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, req)

	if recorder.Code != http.StatusInternalServerError {
		t.Fatalf("Expected status %d, but got %d", http.StatusInternalServerError, recorder.Code)
	}
}
