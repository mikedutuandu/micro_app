package services

import (
	"fmt"
	"github.com/spf13/viper"
	micro_event_pb "github.com/mikedutuandu/micro_app/event_store/protos"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
	"os"
)

var TokenSv TokenService
var EventClient micro_event_pb.EventStoreServiceClient

func Init() {
	//1. Init token service
	TokenSv = TokenService{}
	//2. Init event client service


	var microEventHost = os.Getenv("MICRO_EVENT_HOST")
	if microEventHost == ""{
		microEventHost =  viper.GetString("event.eventHostDefault")
	}
	var microEventPort = os.Getenv("MICRO_EVENT_PORT")
	if microEventPort == "" {
		microEventPort =  viper.GetString("event.eventPortDefault")
	}

	tls := false
	opts := grpc.WithInsecure()
	if tls {
		certFile := "ssl/ca.crt" // Certificate Authority Trust certificate
		creds, sslErr := credentials.NewClientTLSFromFile(certFile, "")
		if sslErr != nil {
			log.Fatalf("Error while loading CA trust certificate: %v", sslErr)
		}
		opts = grpc.WithTransportCredentials(creds)
	}


	ccEvent, err := grpc.Dial(microEventHost+":"+microEventPort, opts)
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}



	eventClient := micro_event_pb.NewEventStoreServiceClient(ccEvent)
	fmt.Printf("Created client: %f", eventClient)


	EventClient = eventClient

	println("Init Service")

}