package repository

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/aabdullahgungor/go-restapi-mock/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	ErrCarNotFound = errors.New("FromRepository - car not found")
)

type MongoDbCarRepository struct {
	connectionPool *mongo.Database
}

func NewMongoDbCarRepository() *MongoDbCarRepository {
	databaseURL := "mongodb+srv://<username>:<password>@cluster0.xbwcqpz.mongodb.net/?retryWrites=true&w=majority"
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()   
    // mongo.Connect return mongo.Client method
    client, err := mongo.Connect(ctx, options.Client().ApplyURI(databaseURL))
    if err != nil {
        panic(err)
    }
    
    db := client.Database("vehicledb")
     
    return &MongoDbCarRepository{
		connectionPool: db,
	}
}

func (m *MongoDbCarRepository) GetAllCars() ([]model.Car, error) {

	carCollection := m.connectionPool.Collection("cars")
	
	var cars []model.Car
	carCursor, err := carCollection.Find(context.TODO(), bson.M{})
	if err != nil {
        panic(err)
	}
	if err = carCursor.All(context.TODO(), &cars); err != nil {
        panic(err)
	}
	
	return cars, err
}

func (m *MongoDbCarRepository) GetCarById(id string) (model.Car, error) { 

	carCollection := m.connectionPool.Collection("cars")

	objId, _ := primitive.ObjectIDFromHex(id)
	
	var car model.Car
	err := carCollection.FindOne(context.TODO(), bson.M{"_id": objId}).Decode(&car)
	if err != nil {
		return model.Car{}, ErrCarNotFound
	}
	
	return car, nil

}

func (m *MongoDbCarRepository) CreateCar(car *model.Car) error {

	carCollection := m.connectionPool.Collection("cars")

	result, err := carCollection.InsertOne(context.TODO(), car)

	if err != nil {
        panic(err)
	}

	log.Printf("\ndisplay the ids of the newly inserted objects: %v", result.InsertedID)

	return  err
}

func (m *MongoDbCarRepository) EditCar(car *model.Car) error { 

	carCollection := m.connectionPool.Collection("car")

	bsonBytes, err:= bson.Marshal(&car)
	
	if err != nil {
            panic(fmt.Errorf("can't marshal:%s", err))
    }

	var upt bson.M
	bson.Unmarshal(bsonBytes, &upt)

	update := bson.M{"$set": upt,}

	result, err := carCollection.UpdateOne(context.TODO(), bson.M{"_id": car.Id}, update)

	if err != nil {
        panic(err)
	}

	log.Println("Number of documents updated:"+ strconv.Itoa(int(result.ModifiedCount))) 

	return  err
}

func (m *MongoDbCarRepository) DeleteCar(id string) error { 

	carCollection := m.connectionPool.Collection("cars")

	objId, _ := primitive.ObjectIDFromHex(id)
	
	result, err := carCollection.DeleteOne(context.TODO(), bson.M{"_id": objId})

	// check for errors in the deleting
	if err != nil {
        panic(err)
	}

	// display the number of documents deleted
	log.Println("deleting the first result from the search filter\n"+ "Number of documents deleted:"+strconv.Itoa(int(result.DeletedCount)))

	return err
}