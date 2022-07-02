package chat

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/dshurubtsov/internal/chat"
	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_Handler_CreateChat(t *testing.T) {
	//Arrange
	router := httprouter.New()
	service := new(ServiceMock)
	service.Mock.On("CreateChat", mock.Anything, mock.Anything).Return(nil)

	handler := chat.NewHandler(service)
	handler.Register(router)

	rr := httptest.NewRecorder()

	// test cases
	tests := []struct {
		name         string
		testChat     chat.Chat
		codeExpected int
	}{
		{
			name: "test 1",
			testChat: chat.Chat{
				Name:            "chat1",
				FounderNickname: "creator1",
			},
			codeExpected: http.StatusCreated,
		},
		{
			name: "test 2",
			testChat: chat.Chat{
				Name:            "chat_test",
				FounderNickname: "test_creator"},
			codeExpected: http.StatusCreated,
		},
	}

	//Act
	for _, tt := range tests {
		body, _ := json.Marshal(tt.testChat)
		request, err := http.NewRequest(http.MethodPost, "/chats/create", bytes.NewBuffer(body))
		assert.NoError(t, err)

		router.ServeHTTP(rr, request)
		assert.Equal(t, tt.codeExpected, rr.Code, fmt.Sprint("Name test: ", tt.name))
	}
}
