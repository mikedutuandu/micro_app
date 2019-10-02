package handlers

import (
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
	"log"
	microbookingpb "github.com/mikedutuandu/micro_app/micro_booking/protos"
	"fmt"
	"os"
)


func Init(){
	var port = os.Getenv("PORT")
	if port == ""{
		port = viper.GetString("gprc.portDefault")
	}
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	opts := []grpc.ServerOption{}
	s := grpc.NewServer(opts...)

	bookingService := &BookingService{}


	microbookingpb.RegisterBookingServiceServer(s, bookingService)
	// Register reflection service on gRPC server.
	reflection.Register(s)

	fmt.Println("Starting Server...")
	if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
}