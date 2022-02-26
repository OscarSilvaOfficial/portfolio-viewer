package db

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Mongo struct {
	stringConnection string
	database         string
	collection       string
	connection       *mongo.Collection
	ctx              context.Context
}

func (m *Mongo) CreateDocument(document map[string]string) {
	m.connection.InsertOne(m.ctx, document)
}

func (m *Mongo) SetStringConnection(stringConnection string) *Mongo {
	m.stringConnection = stringConnection
	return m
}

func (m *Mongo) SetDatabase(database string) *Mongo {
	m.database = database
	return m
}

func (m *Mongo) SetCollection(collection string) *Mongo {
	m.collection = collection
	return m
}

func (m *Mongo) CreateConnection() {

	clientOptions := options.Client().ApplyURI(m.stringConnection)

	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("[Mongo Connection] Connected to MongoDB!")

	m.connection = client.Database(m.database).Collection(m.collection)
	m.ctx = context.TODO()

}

func FactoryMongo(stringConnection string, database string, colleciton string) Mongo {
	mongo := Mongo{}
	mongo.SetStringConnection(stringConnection).SetDatabase(database).SetCollection(colleciton).CreateConnection()
	return mongo
}
