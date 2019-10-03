package services

import (
	micro_teacher_pb "github.com/mikedutuandu/micro_app/micro_teacher/protos"
	micro_learner_pb "github.com/mikedutuandu/micro_app/micro_learner/protos"
)


type MicroClient struct{
	MicroTeacherClient micro_teacher_pb.TeacherServiceClient
	MicroLearnerClient micro_learner_pb.LearnerServiceClient
}

