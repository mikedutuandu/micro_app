package services

import (
	micro_booking_pb "github.com/mikedutuandu/micro_app/micro_booking/protos"
	micro_teacher_pb "github.com/mikedutuandu/micro_app/micro_teacher/protos"
)


type MicroClient struct{
	microTeacherClient micro_teacher_pb.TeacherServiceClient
	microBookingClient micro_booking_pb.BookingServiceClient
}

