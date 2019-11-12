#!/bin/bash

protoc micro_teacher/protos/teacher.proto --go_out=plugins=grpc:.
protoc micro_teacher/protos/auth.proto --go_out=plugins=grpc:.
protoc micro_learner/protos/auth.proto --go_out=plugins=grpc:.
protoc micro_booking/protos/booking.proto --go_out=plugins=grpc:.
protoc micro_learner/protos/learner.proto --go_out=plugins=grpc:.
protoc micro_payment/protos/payment.proto --go_out=plugins=grpc:.
protoc event_store/protos/eventstore.proto --go_out=plugins=grpc:.