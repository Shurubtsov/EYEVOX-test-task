package chat

import (
	"context"

	"github.com/dshurubtsov/internal/chat"
	"github.com/stretchr/testify/mock"
)

type RepositoryMock struct {
	mock.Mock
}

func (m *RepositoryMock) Create(ctx context.Context, chat *chat.Chat) error {
	args := m.Called(ctx, chat)
	return args.Error(0)
}
