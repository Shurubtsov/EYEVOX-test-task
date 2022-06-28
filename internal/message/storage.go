package message

import "context"

type Repository interface {
	Create(ctx context.Context, chat *Message) error
	FindChat(ctx context.Context, name string) error
}
