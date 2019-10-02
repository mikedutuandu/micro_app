package subscribers

import (
	"cloud.google.com/go/pubsub"
	"context"
	"github.com/spf13/viper"

	"google.golang.org/api/option"
)


var BookingSub bookingSubscriber
func Init()  {
	//Init
	ctx := context.Background()
	projectId := viper.GetString("pub.projectID")
	client, err := pubsub.NewClient(ctx, projectId,
		option.WithCredentialsFile("resources/pubsub_service_account.json"),
	)
	if err != nil {
		// TODO: Handle error.
	}
	BookingSub.client = client

	//Subscribe
	go BookingSub.SubHello()
	println("Init Sub")

}