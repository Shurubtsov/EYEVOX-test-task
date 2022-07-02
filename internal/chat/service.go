package chat

import (
	"context"
	"errors"

	"github.com/dshurubtsov/pkg/logging"
)

type ChatService interface {
	CreateChat(ctx context.Context, chat *Chat) error
}

type service struct {
	logger     *logging.Logger
	repository Repository
}

func NewService(rep Repository, logger *logging.Logger) ChatService {
	return &service{
		repository: rep,
		logger:     logger,
	}
}

func (s *service) CreateChat(ctx context.Context, chat *Chat) error {
	if chat.Name == "" || chat.FounderNickname == "" {
		return errors.New("cant' create empty struct")
	}
	// create chat from repository
	if err := s.repository.Create(ctx, chat); err != nil {
		//s.logger.Errorf("Error with create some chat from service, err: %v", err)
		return err
	}
	return nil
}
