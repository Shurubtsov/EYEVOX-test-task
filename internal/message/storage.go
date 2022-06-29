package message

import "context"

type Repository interface {
	Create(ctx context.Context, msg *Message) error
	FindChatID(ctx context.Context, nameChat string) (int, error)
	FindAllByChat(ctx context.Context, chatID, limit, offset int) ([]Message, error)
	FindByID(ctx context.Context, id int, msg *Message) (*Message, error)
}
