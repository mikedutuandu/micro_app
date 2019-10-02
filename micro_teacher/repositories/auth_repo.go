package repositories

import (
	"context"
	micro_teacher_pb "github.com/mikedutuandu/micro_app/micro_teacher/protos"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)



type authRepository struct{
	Client *mongo.Client
	Collection *mongo.Collection
}

func (r authRepository) CreateUser(data *micro_teacher_pb.User) (primitive.ObjectID,error) {
	res, err := r.Collection.InsertOne(context.Background(), data)

	var empty interface{}
	if err != nil {
		return empty.(primitive.ObjectID),err
	}
	oid, ok := res.InsertedID.(primitive.ObjectID)
	if !ok {
		return empty.(primitive.ObjectID),err
	}
	return oid,nil
}

func (r authRepository) GetUserByEmail(email string) (*micro_teacher_pb.User,error) {

	// create an empty struct
	data := &micro_teacher_pb.User{}
	filter := bson.M{"email": email}

	res := r.Collection.FindOne(context.Background(), filter)
	if err := res.Decode(data); err != nil {
		return nil, err
	}
	return data,nil
}

