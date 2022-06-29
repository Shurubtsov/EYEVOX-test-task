package message

import (
	"context"
	"errors"

	"github.com/dshurubtsov/internal/message"
	"github.com/dshurubtsov/pkg/client/postgresql"
	"github.com/dshurubtsov/pkg/logging"
	"github.com/jackc/pgconn"
)

type repository struct {
	client postgresql.Client
	logger *logging.Logger
}

func NewRepository(client postgresql.Client, logger *logging.Logger) message.Repository {
	return &repository{client: client, logger: logger}
}

func (r *repository) Create(ctx context.Context, msg *message.Message) error {
	q := `INSERT INTO messages(creator_nickname, chat_id, text_message) 
		  VALUES ($1, $2, $3) RETURNING id`
	if err := r.client.QueryRow(ctx, q, msg.CreatorNickname, msg.ChatID, msg.TextMessage).Scan(&msg.ID); err != nil {
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

func (r *repository) FindChatID(ctx context.Context, nameChat string) (int, error) {
	var id int
	q := `SELECT id FROM chats WHERE name=$1;`
	if err := r.client.QueryRow(ctx, q, nameChat).Scan(&id); err != nil {
		var pgErr *pgconn.PgError
		if errors.Is(err, pgErr) {
			pgErr = err.(*pgconn.PgError)
			r.logger.Errorf("SQL Error message (%s), Details: %s where -> %s", pgErr.Message, pgErr.Detail, pgErr.Where)
			return 0, nil
		}
		return 0, err
	}

	return id, nil
}

func (r *repository) FindAllByChat(ctx context.Context, chatID, limit, offset int) ([]message.Message, error) {
	q := `SELECT id FROM messages WHERE chat_id=$1 LIMIT $2 OFFSET $3`
	rows, err := r.client.Query(ctx, q, chatID, limit, offset)
	if err != nil {
		r.logger.Error("Error get rows from client Query")
		return nil, err
	}

	msges := make([]message.Message, 0)
	for rows.Next() {
		var msg message.Message
		err = rows.Scan(&msg.ID)
		if err != nil {
			r.logger.Error("Error to scan row")
			return nil, err
		}

		msges = append(msges, msg)
	}
	if err = rows.Err(); err != nil {
		r.logger.Error("[Rows] error")
		return nil, err
	}

	return msges, nil
}

func (r *repository) FindByID(ctx context.Context, id int, msg *message.Message) (*message.Message, error) {
	q := `SELECT chats.name, messages.id, messages.creator_nickname, messages.text_message 
		  FROM chats INNER JOIN messages ON chats.id = messages.chat_id 
		  WHERE messages.id = $1`

	if err := r.client.QueryRow(ctx, q, id).Scan(&msg.ChatName, &msg.ID, &msg.CreatorNickname, &msg.TextMessage); err != nil {
		var pgErr *pgconn.PgError
		if errors.Is(err, pgErr) {
			pgErr = err.(*pgconn.PgError)
			r.logger.Errorf("SQL Error message (%s), Details: %s where -> %s", pgErr.Message, pgErr.Detail, pgErr.Where)
			return nil, nil
		}
		return nil, err
	}

	return msg, nil
}
