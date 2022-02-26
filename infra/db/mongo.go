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

func (self *Mongo) CreateDocument(document map[string]string) {
	self.connection.InsertOne(self.ctx, document)
}

func (self *Mongo) SetStringConnection(stringConnection string) *Mongo {
	self.stringConnection = stringConnection
	return self
}

func (self *Mongo) SetDatabase(database string) *Mongo {
	self.database = database
	return self
}

func (self *Mongo) SetCollection(collection string) *Mongo {
	self.collection = collection
	return self
}

func (self *Mongo) CreateConnection() {

	clientOptions := options.Client().ApplyURI(self.stringConnection)

	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("[Mongo Connection] Connected to MongoDB!")

	self.connection = client.Database(self.database).Collection(self.collection)
	self.ctx = context.TODO()

}

func FactoryMongo(stringConnection string, database string, colleciton string) Mongo {
	mongo := Mongo{}
	mongo.SetStringConnection(stringConnection).SetDatabase(database).SetCollection(colleciton).CreateConnection()
	return mongo
}
