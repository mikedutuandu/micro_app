package repositories

import (
	"context"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)



var BookingRepo bookingRepository
func Init() {

	mongoConnectString := viper.GetString("repo.mongoConnectString")
	mongoDBName := viper.GetString("repo.mongoDBName")

	// if we crash the go code, we get the file name and line number
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	// connect to MongoDB
	client, err := mongo.NewClient(options.Client().ApplyURI(mongoConnectString))
	if err != nil {
		log.Fatal(err)
	}
	err = client.Connect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	BookingRepo.Collection =client.Database(mongoDBName).Collection("booking")
	BookingRepo.Client = client
	println("Init Repo")

}