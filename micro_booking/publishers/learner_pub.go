package publishers

import (
	"cloud.google.com/go/pubsub"
	"context"
	"fmt"
)



type LearnerPublisher struct{
	client *pubsub.Client
}


func (p LearnerPublisher) pubMessage(){
	//test pub
	ctx := context.Background()

	topic := p.client.Topic("topicName")
	ok, err := topic.Exists(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	if !ok {
		// Topic doesn't exist.
		topic, err = p.client.CreateTopic(ctx, "topicName")
		if err != nil {
			// TODO: Handle error.
		}
	}

	defer topic.Stop()
	var results []*pubsub.PublishResult
	r := topic.Publish(ctx, &pubsub.Message{
		Data: []byte("hello world"),
	})
	results = append(results, r)
	// Do other work ...
	for _, r := range results {
		id, err := r.Get(ctx)
		if err != nil {
			// TODO: Handle error.
		}
		fmt.Printf("Published a message with a message ID: %s\n", id)
	}
}
