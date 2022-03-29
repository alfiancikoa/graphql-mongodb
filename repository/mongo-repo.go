package repository

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/alfiancikoa/graphql-mongodb/graph/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MovieRepository interface {
	Save(video *model.Movie)
	GetAll() []*model.Movie
}

type database struct {
	client *mongo.Client
}

var ctx = context.Background()

func New() MovieRepository {

	clientOptions := options.Client().ApplyURI("mongodb://172.17.0.2:27017")
	clientOptions = clientOptions.SetMaxPoolSize(50)

	contex, _ := context.WithTimeout(ctx, 30*time.Second)

	dbClinet, err := mongo.Connect(contex, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("DATABASE is CONNECTED to MongoDB")
	return &database{
		client: dbClinet,
	}
}

const (
	DATABASE   = "graph"
	COLLECTION = "movies"
)

func (db *database) Save(video *model.Movie) {
	collection := db.client.Database(DATABASE).Collection(COLLECTION)
	_, err := collection.InsertOne(context.TODO(), video)
	if err != nil {
		log.Fatal(err)
	}
}

func (db *database) GetAll() []*model.Movie {
	collection := db.client.Database(DATABASE).Collection(COLLECTION)
	cursor, err := collection.Find(context.TODO(), bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(ctx)
	var data []*model.Movie

	for cursor.Next(ctx) {
		var v *model.Movie
		err := cursor.Decode(&v)
		if err != nil {
			log.Fatal(err)
		}
		data = append(data, v)
	}

	return data
}
