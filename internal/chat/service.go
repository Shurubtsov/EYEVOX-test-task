package chat

import (
	"context"

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
	if err := s.repository.Create(ctx, chat); err != nil {
		s.logger.Errorf("Error with create some chat from service, err: %v", err)
		return err
	}
	return nil
}
