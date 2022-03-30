package dreampubsub

import (
	"context"

	"cloud.google.com/go/pubsub"
	"github.com/andreashanson/golang-pusub-cloud-function/pkg/message"
)

type msgID string

type PubSubRepository struct {
	client *pubsub.Client
	ctx    context.Context
}

func NewPubSubRepository(projectID string) (*PubSubRepository, error) {
	ctx := context.Background()

	client, err := pubsub.NewClient(ctx, projectID)
	if err != nil {
		return &PubSubRepository{}, err
	}

	return &PubSubRepository{client: client}, nil
}

func (pbr *PubSubRepository) Publish(topic string, m string) (message.Message, error) {
	msg := pubsub.Message{
		Data: []byte(m),
	}

	t := pbr.client.Topic(topic)
	res := t.Publish(pbr.ctx, &msg)

	mid, err := res.Get(pbr.ctx)
	if err != nil {
		return message.Message{}, err
	}

	mm := message.Message{ID: mid, Data: m}

	return mm, nil
}
