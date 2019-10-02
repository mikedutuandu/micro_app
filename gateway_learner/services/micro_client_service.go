package services

import (
	micro_booking_pb "github.com/mikedutuandu/micro_app/micro_booking/protos"
	micro_learner_pb "github.com/mikedutuandu/micro_app/micro_learner/protos"
	micro_teacher_pb "github.com/mikedutuandu/micro_app/micro_teacher/protos"
)


type MicroClient struct{
	MicroLearnerClient micro_learner_pb.LearnerServiceClient
	MicroBookingClient micro_booking_pb.BookingServiceClient
	MicroTeacherClient micro_teacher_pb.AuthServiceClient
}



