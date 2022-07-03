package message

import (
	"context"
	"testing"

	"github.com/dshurubtsov/internal/message"
	"github.com/dshurubtsov/pkg/logging"
	"github.com/stretchr/testify/assert"
)

func TestService_CreateMessage(t *testing.T) {
	type mockReturn func(r *RepostioryMock, inputMessage message.Message, chatName string)

	tests := []struct {
		name         string
		chatName     string
		inputMessage message.Message
		mockReturn   mockReturn
		wantErr      error
	}{
		{
			name:     "input OK",
			chatName: "test",
			inputMessage: message.Message{
				ChatID:          1,
				ChatName:        "test",
				CreatorNickname: "test",
				TextMessage:     "test test test",
			},
			mockReturn: func(r *RepostioryMock, inputMessage message.Message, chatName string) {
				r.Mock.On("Create", &inputMessage).Return(nil)
				r.Mock.On("FindChatID", chatName).Return(1, nil)
			},
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRep := new(RepostioryMock)
			logger := logging.GetLogger()
			service := message.NewService(mockRep, logger)

			tt.mockReturn(mockRep, tt.inputMessage, tt.chatName)

			err := service.CreateMessage(context.TODO(), &tt.inputMessage, tt.chatName)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}
