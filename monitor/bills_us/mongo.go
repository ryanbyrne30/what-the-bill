package bills_us

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type BillActions struct {
	Date   string `bson:"date"`
	Action string `bson:"action"`
}

type BillMembers struct {
	Role    string `bson:"role"`
	Chamber string `bson:"chamber"`
	BioID   string `bson:"bio_id"`
	Name    string `bson:"name"`
	State   string `bson:"state"`
	Party   string `bson:"party"`
}

type BillCommittees struct {
	AuthorityID string `bson:"authority_id"`
	Chamber     string `bson:"chamber"`
	Name        string `bson:"name"`
	Type        string `bson:"type"`
}

type Bill struct {
	ID          primitive.ObjectID `bson:"_id"`
	BillID      string             `bson:"bill_id"`
	ShortTitle  string             `bson:"short_title"`
	Title       string             `bson:"title"`
	Url         string             `bson:"url"`
	Text        string             `bson:"text"`
	Congress    string             `bson:"congress"`
	Session     string             `bson:"session"`
	Pages       int                `bson:"pages"`
	Version     string             `bson:"version"`
	Type        string             `bson:"type"`
	Members     []BillMembers      `bson:"members"`
	Committees  []BillCommittees   `bson:"committees"`
	Actions     []BillActions      `bson:"actions"`
	PublishedAt time.Time          `bson:"published_at"`
	UpdatedAt   time.Time          `bson:"updated_at"`
}

type Mongo struct {
	client *mongo.Client
	col    *mongo.Collection
}

func NewMongo(url, db string, col string) *Mongo {
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

func (m *Mongo) InsertBill(ctx context.Context, bill *Bill) error {
	_, err := m.col.InsertOne(ctx, bill)
	return err
}
