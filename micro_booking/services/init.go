package services

import (
	micro_teacher_pb "github.com/mikedutuandu/micro_app/micro_teacher/protos"
	micro_learner_pb "github.com/mikedutuandu/micro_app/micro_learner/protos"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"os"
	"log"
	"fmt"
)

var MicroCLI MicroClient
func Init() {

	//1. Init micro client service
	var microLearnerHost = os.Getenv("MICRO_LEARNER_HOST")
	if microLearnerHost == ""{
		microLearnerHost = viper.GetString("micro.learnerHostDefault")
	}
	var microLearnerPort = os.Getenv("MICRO_LEARNER_PORT")
	if microLearnerPort == "" {
		microLearnerPort = viper.GetString("micro.learnerPortDefault")
	}

	var microBookingHost = os.Getenv("MICRO_BOOKING_HOST")
	if microBookingHost == ""{
		microBookingHost = viper.GetString("micro.bookingHostDefault")
	}
	var microBookingPort = os.Getenv("MICRO_BOOKING_PORT")
	if microBookingPort == "" {
		microBookingPort = viper.GetString("micro.bookingPortDefault")
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

	ccLearner, err := grpc.Dial(microLearnerHost+":"+microLearnerPort, opts)
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	//defer ccLearner.Close()


	ccBooking, err := grpc.Dial(microBookingHost+":"+microBookingPort, opts)
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	//defer ccBooking.Close()


	learnerClient := micro_learner_pb.NewLearnerServiceClient(ccLearner)
	fmt.Printf("Created client: %f", learnerClient)

	bookingClient := micro_teacher_pb.NewAuthServiceClient(ccBooking)
	fmt.Printf("Created client: %f", bookingClient)


	MicroCLI.microLearnerClient = learnerClient
	MicroCLI.microTeacherClient = bookingClient

	println("Init Service")

}