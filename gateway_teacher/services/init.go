package services

import (
	"fmt"
	micro_booking_pb "github.com/mikedutuandu/micro_app/micro_booking/protos"
	micro_learner_pb "github.com/mikedutuandu/micro_app/micro_learner/protos"
	micro_payment_pb "github.com/mikedutuandu/micro_app/micro_payment/protos"
	micro_teacher_pb "github.com/mikedutuandu/micro_app/micro_teacher/protos"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
	"os"
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
		microBookingHost =  viper.GetString("micro.bookingHostDefault")
	}
	var microBookingPort = os.Getenv("MICRO_BOOKING_PORT")
	if microBookingPort == "" {
		microBookingPort = viper.GetString("micro.bookingPortDefault")
	}

	var microTeacherHost = os.Getenv("MICRO_TEACHER_HOST")
	if microTeacherHost == ""{
		microTeacherHost =  viper.GetString("micro.teacherHostDefault")
	}
	var microTeacherPort = os.Getenv("MICRO_TEACHER_PORT")
	if microTeacherPort == "" {
		microTeacherPort =  viper.GetString("micro.teacherPortDefault")
	}

	var microPaymentHost = os.Getenv("MICRO_PAYMENT_HOST")
	if microPaymentHost == ""{
		microPaymentHost =  viper.GetString("micro.paymentHostDefault")
	}
	var microPaymentPort = os.Getenv("MICRO_PAYMENT_PORT")
	if microPaymentPort == "" {
		microPaymentPort =  viper.GetString("micro.paymentPortDefault")
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


	ccTeacher, err := grpc.Dial(microTeacherHost+":"+microTeacherPort, opts)
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	//defer ccTeacher.Close()

	ccPayment, err := grpc.Dial(microPaymentHost+":"+microPaymentPort, opts)
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}


	learnerClient := micro_learner_pb.NewLearnerServiceClient(ccLearner)
	fmt.Printf("Created client: %f", learnerClient)

	bookingClient := micro_booking_pb.NewBookingServiceClient(ccBooking)
	fmt.Printf("Created client: %f", bookingClient)

	teacherAuthClient := micro_teacher_pb.NewAuthServiceClient(ccTeacher)
	fmt.Printf("Created client: %f", teacherAuthClient)

	teacherClient := micro_teacher_pb.NewTeacherServiceClient(ccTeacher)
	fmt.Printf("Created client: %f", teacherClient)

	paymentClient := micro_payment_pb.NewPaymentServiceClient(ccPayment)
	fmt.Printf("Created client: %f", bookingClient)

	MicroCLI.MicroLearnerClient = learnerClient
	MicroCLI.MicroBookingClient = bookingClient
	MicroCLI.MicroTeacherClient = teacherClient
	MicroCLI.MicroTeacherAuthClient = teacherAuthClient
	MicroCLI.MicroPaymentClient = paymentClient

	println("Init Service")

}