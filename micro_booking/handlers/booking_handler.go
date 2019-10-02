package handlers

import (
	"context"
	"log"
	//publishers "github.com/mikedutuandu/micro_app/micro_learner/publishers"

	microbookingpb "github.com/mikedutuandu/micro_app/micro_booking/protos"
)

type BookingService struct {

}



func (srv *BookingService) Hello(ctx context.Context, req *microbookingpb.HelloRequest) (*microbookingpb.HelloResponse, error) {
	log.Println("go into hello")

	return &microbookingpb.HelloResponse{
	}, nil
}
