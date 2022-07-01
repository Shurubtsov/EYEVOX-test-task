package chat

import (
	"context"
	"testing"
	"time"

	chatPkg "github.com/dshurubtsov/internal/chat"
	"github.com/dshurubtsov/pkg/logging"
	"github.com/stretchr/testify/mock"
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

	cht := chatPkg.Chat{
		Name:            "bullychat",
		FounderNickname: "korki",
	}
	cht2 := chatPkg.Chat{
		Name:            "",
		FounderNickname: "_",
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
				chat: &cht,
			},
			wantErr: false,
		},
		{
			name: "test2",
			s:    service,
			args: args{
				ctx:  ctx,
				chat: &cht2,
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRep.On("Create", mock.Anything).Return(nil)
			if err := tt.s.CreateChat(tt.args.ctx, tt.args.chat); (err != nil) != tt.wantErr {
				t.Errorf("service.CreateChat() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
