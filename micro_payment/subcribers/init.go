package subscribers

import (
	"cloud.google.com/go/pubsub"
	"context"
	"github.com/spf13/viper"

	"google.golang.org/api/option"
)


var EventSub eventSubscriber
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
	EventSub.client = client

	//Subscribe
	go EventSub.SubHello()
	println("Init Sub")

}