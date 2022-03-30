package producer

import "github.com/andreashanson/golang-pusub-cloud-function/pkg/message"

type Repository interface {
	Publish(topic string, msg string) (message.Message, error)
}

type Service struct {
	repo Repository
}

func NewService(r Repository) *Service {
	return &Service{
		repo: r,
	}
}

func (s *Service) Publish(topic string, msg string) (message.Message, error) {
	return s.repo.Publish(topic, msg)
}
