package main

import (
	"fmt"

	"github.com/andreashanson/golang-pusub-cloud-function/pkg/external/dreampubsub"
)

func main() {
	pubsub := dreampubsub.NewPubSub("randomNumbers")
	fmt.Println(pubsub.Topic)
}
