package chat

import (
	"context"
	"errors"
	"testing"

	"github.com/dshurubtsov/internal/chat"
	//chatPkg "github.com/dshurubtsov/internal/chat"
	"github.com/dshurubtsov/pkg/logging"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_service_CreateChat(t *testing.T) {

	// mock behavior with return others
	type mockReturn func(s *RepositoryMock, inputChat chat.Chat)

	// test cases
	tests := []struct {
		name       string
		inputChat  chat.Chat
		mockReturn mockReturn
		wantErr    error
	}{
		{
			name: "input OK",
			inputChat: chat.Chat{
				Name:            "test",
				FounderNickname: "test",
			},
			mockReturn: func(s *RepositoryMock, inputChat chat.Chat) {
				s.Mock.On("Create", mock.Anything, &inputChat).Return(nil)
			},
			wantErr: nil,
		},
		{
			name: "input BAD",
			inputChat: chat.Chat{
				Name:            "",
				FounderNickname: "",
			},
			mockReturn: func(s *RepositoryMock, inputChat chat.Chat) {
				s.Mock.On("Create", mock.Anything, &inputChat).Return(nil)
			},
			wantErr: errors.New("cant' create empty struct"),
		},
	}

	// performance our tests
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// init dependencies
			mockRep := new(RepositoryMock)
			logger := logging.GetLogger()
			service := chat.NewService(mockRep, logger)

			// setup mock calls
			tt.mockReturn(mockRep, tt.inputChat)

			// execute method from mock and assert test
			err := service.CreateChat(context.TODO(), &tt.inputChat)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}
