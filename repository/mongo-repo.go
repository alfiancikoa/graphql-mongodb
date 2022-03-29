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
	Save(video *model.Movie) error
	GetAll() ([]*model.Movie, error)
	Get(id int) (*model.Movie, error)
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

// Query database untuk menyimpan data ke dalam database Mongodb
func (db *database) Save(video *model.Movie) error {
	collection := db.client.Database(DATABASE).Collection(COLLECTION)
	_, err := collection.InsertOne(context.TODO(), video)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

// Query database untuk mendapatkan data berdasarkan input id
func (db *database) Get(id int) (*model.Movie, error) {
	collection := db.client.Database(DATABASE).Collection(COLLECTION)
	cursor, err := collection.Find(ctx, bson.M{"id": id})
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	var data *model.Movie
	for cursor.Next(ctx) {
		err := cursor.Decode(&data)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
	}
	return data, nil
}

// Query database untuk mendapatkan seluruh data pada database
func (db *database) GetAll() ([]*model.Movie, error) {
	collection := db.client.Database(DATABASE).Collection(COLLECTION)
	cursor, err := collection.Find(context.TODO(), bson.D{})
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer cursor.Close(ctx)
	var data []*model.Movie

	for cursor.Next(ctx) {
		var v *model.Movie
		err := cursor.Decode(&v)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		data = append(data, v)
	}

	return data, nil
}
