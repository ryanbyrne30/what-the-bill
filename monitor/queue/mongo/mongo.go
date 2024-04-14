package mongo

import (
	"context"
	"errors"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Event struct {
	ID        primitive.ObjectID `bson:"_id"`
	EventName string             `bson:"event_name"`
	Message   string             `bson:"message"`
	Created   time.Time          `bson:"created_at"`
}

type Mongo struct {
	client *mongo.Client
	col    *mongo.Collection
}

func NewMongo(url, db, col string) *Mongo {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(url))
	if err != nil {
		log.Fatalf("could not connect to Mongo: %v", err)
	}
	log.Printf("successfully connected to Mongo")

	coll := client.Database(db).Collection(col)
	return &Mongo{client: client, col: coll}
}

func (m *Mongo) Disconnect(ctx context.Context) error {
	return m.client.Disconnect(ctx)
}

func (m *Mongo) SaveEvent(ctx context.Context, event *Event) error {
	if event == nil {
		return errors.New("no event given")
	}
	event.ID = primitive.NewObjectID()
	event.Created = time.Now()
	_, err := m.col.InsertOne(ctx, event)
	return err
}
