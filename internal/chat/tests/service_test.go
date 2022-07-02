package chat

import (
	"context"
	"errors"
	"testing"
	"time"

	chatPkg "github.com/dshurubtsov/internal/chat"
	"github.com/dshurubtsov/pkg/logging"
)

func Test_service_CreateChat(t *testing.T) {
	type args struct {
		ctx  context.Context
		chat *chatPkg.Chat
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	mockRep := new(RepositoryMock)
	logger := logging.GetLogger()
	service := chatPkg.NewService(mockRep, logger)

	testEntities := []chatPkg.Chat{
		{
			Name:            "bullychat",
			FounderNickname: "korki",
		},
		{
			Name:            "",
			FounderNickname: "_",
		},
		{
			Name:            "String",
			FounderNickname: "",
		},
		{
			Name:            "ErrorDBchat",
			FounderNickname: "error",
		},
	}

	tests := []struct {
		name    string
		s       chatPkg.ChatService
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "test",
			s:    service,
			args: args{
				ctx:  ctx,
				chat: &testEntities[0],
			},
			wantErr: true,
		},
		{
			name: "test2",
			s:    service,
			args: args{
				ctx:  ctx,
				chat: &testEntities[1],
			},
			wantErr: true,
		},
		{
			name: "test3",
			s:    service,
			args: args{
				ctx:  ctx,
				chat: &testEntities[2],
			},
			wantErr: true,
		},
		{
			name: "test4",
			s:    service,
			args: args{
				ctx:  ctx,
				chat: &testEntities[3],
			},
			wantErr: true,
		},
	}

	mockError := errors.New("failed")

	//tests with mock return Error from repository
	for i, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRep.On("Create", ctx, &testEntities[i]).Return(mockError)
			err := tt.s.CreateChat(tt.args.ctx, tt.args.chat)
			if (err != nil) != tt.wantErr {
				t.Errorf("service.CreateChat() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}

	// tests with successfuly return nil from mockRepository
	for i, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRep.On("Create", ctx, &testEntities[i]).Return(mockError)
			if err := tt.s.CreateChat(tt.args.ctx, tt.args.chat); (err != nil) != tt.wantErr {
				t.Errorf("service.CreateChat() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
