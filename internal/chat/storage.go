package chat

import "context"

type Repository interface {
	Create(ctx context.Context, chat *Chat) error
}
