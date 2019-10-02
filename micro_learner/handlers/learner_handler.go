package handlers

import (
	"context"
	"log"
	//publishers "github.com/mikedutuandu/micro_app/micro_learner/publishers"

	microlearnerpb "github.com/mikedutuandu/micro_app/micro_learner/protos"
)

type LearnerService struct {

}



func (srv *LearnerService) Hello(ctx context.Context, req *microlearnerpb.HelloRequest) (*microlearnerpb.HelloResponse, error) {
	log.Println("go into hello")

	return &microlearnerpb.HelloResponse{
	}, nil
}
