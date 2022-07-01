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
	args := m.Called(chat)
	var r0 error
	v1 := args.Get(0)
	if v1 != nil {
		r0 = v1.(error)
	}
	return r0
}
