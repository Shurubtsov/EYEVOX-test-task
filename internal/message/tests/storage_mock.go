package message

import (
	"context"

	"github.com/dshurubtsov/internal/message"
	"github.com/stretchr/testify/mock"
)

type RepostioryMock struct {
	mock.Mock
}

func (r *RepostioryMock) Create(ctx context.Context, msg *message.Message) error {
	args := r.Mock.Called(msg)
	return args.Error(0)
}
func (r *RepostioryMock) FindChatID(ctx context.Context, nameChat string) (int, error) {
	args := r.Mock.Called(nameChat)
	return args.Int(0), args.Error(1)
}
func (r *RepostioryMock) FindAllByChat(ctx context.Context, chatID, limit, offset int) ([]message.Message, error) {
	args := r.Mock.Called(chatID, limit, offset)
	return args.Get(0).([]message.Message), args.Error(1)
}
func (r *RepostioryMock) FindByID(ctx context.Context, id int, msg *message.Message) (*message.Message, error) {
	args := r.Mock.Called(id, msg)
	return args.Get(0).(*message.Message), args.Error(1)
}
