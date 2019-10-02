package handlers

import (
	"context"
	"log"
	//publishers "github.com/mikedutuandu/micro_app/micro_teacher/publishers"

	microteacherpb "github.com/mikedutuandu/micro_app/micro_teacher/protos"
)

type TeacherService struct {

}



func (srv *TeacherService) Hello(ctx context.Context, req *microteacherpb.HelloRequest) (*microteacherpb.HelloResponse, error) {
	log.Println("go into hello")

	return &microteacherpb.HelloResponse{
	}, nil
}
