package database

import (
	"context"
	"github.com/TheChaseBL/pokerepo-backend/utils"
	"github.com/mongodb/mongo-go-driver/mongo"
	"log"
	"github.com/TheChaseBL/pokerepo-backend/entity"
)

type connection struct {
	Client     *mongo.Client
	collection *mongo.Collection
}

var Conn *connection

func NewConnection() *connection {
	config := utils.GetConfig()
	client, err := mongo.NewClient(config.MongoUrl)
	if err != nil { log.Fatal(err) }
	err = client.Connect(context.TODO())
	if err != nil { log.Fatal(err) }

	collection := client.Database("pokerepo").Collection("pokemon")

	newConnect := connection{Client: client, collection: collection}
	return &newConnect
}

func (db connection) AddEntry(p entity.Pokemon) {
	_, err := db.collection.InsertOne(context.Background(), p)
	if err != nil { log.Fatal(err) }
}

func (db connection) ReadAll() []entity.Pokemon {
	pokemon := make([]entity.Pokemon, 0)
	cur, err := db.collection.Find(context.Background(), nil)
	if err != nil { log.Fatal(err) }
	for cur.Next(context.Background()) {
		ele := &entity.Pokemon{}
		cur.Decode(ele)
		pokemon = append(pokemon, *ele)
	}
	return pokemon
}