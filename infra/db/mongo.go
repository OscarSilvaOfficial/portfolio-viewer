package db

type Mongo struct {
	stringConnection string
}

func (m *Mongo) CreateDocument(document map[string]string) map[string]string {
	return document
}

func (m *Mongo) SetStringConnection(stringConnection string) Mongo {
	m.stringConnection = stringConnection
	return *m
}

func FactoryMongo(stringConnection string) Mongo {
	mongo := Mongo{}
	mongo.SetStringConnection(stringConnection)
	return mongo
}