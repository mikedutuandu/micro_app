package services

import (
	micro_booking_pb "github.com/mikedutuandu/micro_app/micro_booking/protos"
	micro_teacher_pb "github.com/mikedutuandu/micro_app/micro_teacher/protos"
)


type MicroClient struct{
	MicroTeacherClient micro_teacher_pb.TeacherServiceClient
	MicroBookingClient micro_booking_pb.BookingServiceClient
}

