package chat

import (
	"context"

	"github.com/dshurubtsov/internal/chat"
	"github.com/stretchr/testify/mock"
)

type ServiceMock struct {
	mock.Mock
}

func (s *ServiceMock) CreateChat(ctx context.Context, chat *chat.Chat) error {
	args := s.Mock.Called(ctx, chat)
	return args.Error(0)
}
