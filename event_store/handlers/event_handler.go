package handlers

import (

	"context"
	//"golang.org/x/crypto/bcrypt"
	//"google.golang.org/grpc/codes"
	//"google.golang.org/grpc/status"
	//"log"
	//
	//
	//"fmt"
	eventstorerpb "github.com/mikedutuandu/micro_app/event_store/protos"

)

type EventStoreService struct {

}



func (srv *EventStoreService) CreateEvent(ctx context.Context, req *eventstorerpb.CreateEventRequest) (*eventstorerpb.CreateEventResponse, error) {

	return &eventstorerpb.CreateEventResponse{
		IsSuccess:true,
	}, nil
}

func (srv *EventStoreService) GetEvents(ctx context.Context, req *eventstorerpb.GetEventsRequest) (*eventstorerpb.GetEventsResponse, error) {

	return &eventstorerpb.GetEventsResponse{
		Events:[]*eventstorerpb.Event{},
	}, nil
}

