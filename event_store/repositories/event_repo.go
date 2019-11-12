package repositories

import (
	"context"
	eventstorerpb "github.com/mikedutuandu/micro_app/event_store/protos"
	//"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)



type eventStoreRepository struct{
	Client *mongo.Client
	Collection *mongo.Collection
}


func (r eventStoreRepository) CreateEvent(event *eventstorerpb.Event) (primitive.ObjectID,error) {
	// Insert two rows into the "accounts" table.
	// sql := fmt.Sprintf("INSERT INTO events (id, eventtype, aggregateid, aggregatetype, eventdata, channel)
	// VALUES ('%s','%s','%s','%s','%s','%s')", event.EventId, event.EventType, event.AggregateId, event.AggregateType, event.EventData, event.Channel)

		res, err := r.Collection.InsertOne(context.Background(), event)

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

func (r eventStoreRepository) GetEvents(filter *eventstorerpb.GetEventsRequest) []*eventstorerpb.Event {
	var events []*eventstorerpb.Event
	return events
}

