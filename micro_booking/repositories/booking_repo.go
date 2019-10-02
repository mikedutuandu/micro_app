package repositories

import (
	"go.mongodb.org/mongo-driver/mongo"
)



type bookingRepository struct{
	Client *mongo.Client
	Collection *mongo.Collection
}


