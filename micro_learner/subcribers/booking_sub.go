package subscribers

import (
	"cloud.google.com/go/pubsub"
	"context"
	"fmt"
	"log"
)



type bookingSubscriber struct{
	client *pubsub.Client
}


func (s bookingSubscriber) SubHello()  {
	ctx := context.Background()

	topic := s.client.Topic("topicName")

	sub := s.client.Subscription("subName")
	ok, err := sub.Exists(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	if !ok {
		// Subscription doesn't exist.
		sub, err = s.client.CreateSubscription(ctx, "subName", pubsub.SubscriptionConfig{
			Topic:            topic,
			//AckDeadline:      10 * time.Second,
			//ExpirationPolicy: 25 * time.Hour,
		})
		if err != nil {
			// TODO: Handle error.
		}
	}

	println("Init Sub")
	// Use a callback to receive messages via subscription.
	err = sub.Receive(ctx, func(ctx context.Context, m *pubsub.Message) {
		fmt.Println(m.Data)
		m.Ack() // Acknowledge that we've consumed the message.
	})
	if err != nil {
		log.Println(err)
	}
}