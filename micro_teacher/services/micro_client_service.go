package services

import (
	micro_booking_pb "github.com/mikedutuandu/micro_app/micro_booking/protos"
	micro_learner_pb "github.com/mikedutuandu/micro_app/micro_learner/protos"
)


type MicroClient struct{
	microLearnerClient micro_learner_pb.LearnerServiceClient
	microBookingClient micro_booking_pb.BookingServiceClient
}

