package message

import (
	"context"

	"github.com/dshurubtsov/pkg/logging"
)

type MessageService interface {
	CreateMessage(ctx context.Context, msg *Message, chatName string) error
	FindMessageByID(ctx context.Context, id int, msg *Message) (*Message, error)
	FindListID(ctx context.Context, chatName string, limit, offset int) ([]Message, error)
}

type service struct {
	logger     *logging.Logger
	repository Repository
}

func NewService(rep Repository, logger *logging.Logger) MessageService {
	return &service{
		repository: rep,
		logger:     logger,
	}
}

func (s *service) CreateMessage(ctx context.Context, msg *Message, chatName string) error {
	var err error
	msg.ChatID, err = s.repository.FindChatID(ctx, chatName)
	if err != nil {
		s.logger.Error(err)
		return err
	}

	err = s.repository.Create(ctx, msg)
	if err != nil {
		s.logger.Error(err)
		return err
	}

	return nil
}

func (s *service) FindListID(ctx context.Context, chatName string, limit, offset int) ([]Message, error) {

	chatID, err := s.repository.FindChatID(ctx, chatName)
	if err != nil {
		s.logger.Error(err)
		return nil, err
	}

	msges, err := s.repository.FindAllByChat(ctx, chatID, limit, offset)
	if err != nil {
		s.logger.Error(err)
		return nil, err
	}

	return msges, nil
}

func (s *service) FindMessageByID(ctx context.Context, id int, msg *Message) (*Message, error) {
	msg, err := s.repository.FindByID(ctx, id, msg)
	if err != nil {
		s.logger.Error("Error find message by id, error: ", err)
		return nil, err
	}

	return msg, nil
}
