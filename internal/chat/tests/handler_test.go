package chat

import (
	"bytes"
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/dshurubtsov/internal/chat"
	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_Handler_CreateChat(t *testing.T) {

	// mock behavior when we call method
	type mockReturn func(m *ServiceMock, inputChat chat.Chat)

	// test cases
	tests := []struct {
		name         string
		inputBody    string
		inputChat    chat.Chat
		mockReturn   mockReturn
		codeExpected int
	}{
		{
			name:      "input OK",
			inputBody: `{"name": "test", "founder_nickname": "test"}`,
			inputChat: chat.Chat{
				Name:            "test",
				FounderNickname: "test",
			},
			mockReturn: func(m *ServiceMock, inputChat chat.Chat) {
				m.Mock.On("CreateChat", mock.Anything, &inputChat).Return(nil)
			},
			codeExpected: http.StatusCreated,
		},
		{
			name:      "input BAD",
			inputBody: `{"name": "", "founder_nickname": ""}`,
			inputChat: chat.Chat{
				Name:            "",
				FounderNickname: "",
			},
			mockReturn: func(m *ServiceMock, inputChat chat.Chat) {
				m.Mock.On("CreateChat", mock.Anything, &inputChat).Return(errors.New("failed"))
			},
			codeExpected: http.StatusInternalServerError,
		},
	}

	// Performance our test cases
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Init deps
			router := httprouter.New()
			service := new(ServiceMock)

			tt.mockReturn(service, tt.inputChat)

			handler := chat.NewHandler(service, context.TODO())
			handler.Register(router)

			// Test Request
			rr := httptest.NewRecorder()
			request, err := http.NewRequest(http.MethodPost, "/chats/create", bytes.NewBuffer([]byte(tt.inputBody)))
			assert.NoError(t, err)

			router.ServeHTTP(rr, request)
			// Asserting
			assert.Equal(t, tt.codeExpected, rr.Code)
		})
	}
}
