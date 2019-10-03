package services

import (
	micro_booking_pb "github.com/mikedutuandu/micro_app/micro_booking/protos"
	micro_teacher_pb "github.com/mikedutuandu/micro_app/micro_teacher/protos"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"os"
	"log"
	"fmt"
)

var TokenSv TokenService
var MicroCLI MicroClient
func Init() {
	//1. Init token service
	TokenSv = TokenService{}

	//2. Init micro client service
	var microTeacherHost = os.Getenv("MICRO_TEACHER_HOST")
	if microTeacherHost == ""{
		microTeacherHost = viper.GetString("micro.teacherHostDefault")
	}
	var microTeacherPort = os.Getenv("MICRO_TEACHER_PORT")
	if microTeacherPort == "" {
		microTeacherPort = viper.GetString("micro.teacherPortDefault")
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

	ccTeacher, err := grpc.Dial(microTeacherHost+":"+microTeacherPort, opts)
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	//defer ccLearner.Close()


	ccBooking, err := grpc.Dial(microBookingHost+":"+microBookingPort, opts)
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	//defer ccBooking.Close()


	teacherClient := micro_teacher_pb.NewTeacherServiceClient(ccTeacher)
	fmt.Printf("Created client: %f", teacherClient)

	bookingClient := micro_booking_pb.NewBookingServiceClient(ccBooking)
	fmt.Printf("Created client: %f", bookingClient)


	MicroCLI.MicroTeacherClient = teacherClient
	MicroCLI.MicroBookingClient = bookingClient

	println("Init Service")

}