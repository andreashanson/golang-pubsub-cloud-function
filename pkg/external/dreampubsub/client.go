package dreampubsub

import (
	"context"

	"cloud.google.com/go/pubsub"
)

type PubSubClient struct {
	Client *pubsub.Client
}

func NewClient(projectID string) (PubSubClient, error) {
	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, projectID)
	if err != nil {
		return PubSubClient{}, err
	}
	return PubSubClient{Client: client}, nil
}
