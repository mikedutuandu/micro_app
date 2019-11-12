package repositories

import (
	"go.mongodb.org/mongo-driver/mongo"
)



type paymentRepository struct{
	Client *mongo.Client
	Collection *mongo.Collection
}


