package chat

import (
	"context"
	"errors"
	"time"

	"github.com/dshurubtsov/internal/chat"
	"github.com/dshurubtsov/pkg/client/postgresql"
	"github.com/dshurubtsov/pkg/logging"
	"github.com/jackc/pgconn"
)

type repository struct {
	client postgresql.Client
	logger *logging.Logger
}

func NewRepository(client postgresql.Client, logger *logging.Logger) chat.Repository {
	return &repository{client: client, logger: logger}
}

func (r *repository) Create(ctx context.Context, chat *chat.Chat) error {

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	q := `INSERT INTO chats(name, founder_nickname) 
		  VALUES ($1, $2) RETURNING id`
	if err := r.client.QueryRow(ctx, q, chat.Name, chat.FounderNickname).Scan(&chat.ID); err != nil {
		var pgErr *pgconn.PgError
		if errors.Is(err, pgErr) {
			pgErr = err.(*pgconn.PgError)
			r.logger.Errorf("SQL Error message (%s), Details: %s where -> %s", pgErr.Message, pgErr.Detail, pgErr.Where)
			return nil
		}
		return err
	}

	return nil
}
