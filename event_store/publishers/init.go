package publishers

import (
	"cloud.google.com/go/pubsub"
	"github.com/spf13/viper"
	"google.golang.org/api/option"
	"context"
)

var Pub Publisher
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
	Pub.client = client
	println("Init Pub")
}