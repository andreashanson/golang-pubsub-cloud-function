package api

import (
	"encoding/json"
	"net/http"

	"github.com/andreashanson/golang-pusub-cloud-function/pkg/external/dreampubsub"
	"github.com/andreashanson/golang-pusub-cloud-function/pkg/producer"
)

var (
	projectID = "amplified-way-344315"
	topic     = "randomNumbers"
)

func PublishOnPubSub(w http.ResponseWriter, r *http.Request) {
	var p struct {
		Message string `json:"message"`
	}
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	pubsubRepo, err := dreampubsub.NewPubSubRepository(projectID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	prod := producer.NewService(pubsubRepo)
	message, err := prod.Publish(topic, p.Message)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(message.Data))

}
