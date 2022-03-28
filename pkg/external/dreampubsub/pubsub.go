package dreampubsub

import (
	"context"

	"cloud.google.com/go/pubsub"
)

type PubSubRepository struct {
	client *pubsub.Client
	ctx    context.Context
}

func NewPubSubRepository(projectID string) (PubSubRepository, error) {
	ctx := context.Background()

	client, err := pubsub.NewClient(ctx, projectID)
	if err != nil {
		return PubSubRepository{}, err
	}
	return PubSubRepository{
		client: client,
	}, nil
}

func (pbr PubSubRepository) Publish(topic string, msg PubSubMessage) PubSubResponse {
	t := pbr.client.Topic(topic)
	return t.Publish(pbr.ctx, msg)
}

type PubSubMessage *pubsub.Message

type PubSubResponse *pubsub.PublishResult
