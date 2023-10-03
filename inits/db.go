package inits

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
	"plants/models"
	"time"
)

var realClient *mongo.Client

func DBInit() {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(os.Getenv("DB_URL")).SetServerAPIOptions(serverAPI)
	// Create a new client and connect to the server
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		panic(err)
	}
	// Send a ping to confirm a successful connection
	if err := client.Database("admin").RunCommand(context.TODO(), bson.D{{"ping", 1}}).Err(); err != nil {
		panic(err)
	}
	fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")
	realClient = client
}

func InsertIntoColletion(collectionName string, plant models.Plant) (*models.Plant, error) {
	database := realClient.Database("test")
	collection := database.Collection(collectionName)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	insertResult, err := collection.InsertOne(ctx, plant)
	cancel()

	if err != nil {
		return nil, err
	}
	plant.ID = insertResult.InsertedID.(primitive.ObjectID)

	return &plant, nil
}

func GetAllInCollection(collectionName string) ([]models.Plant, error) {
	database := realClient.Database("test")
	collection := database.Collection(collectionName)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	var plants []models.Plant
	cursor, err := collection.Find(ctx, bson.D{})
	cancel()

	if err != nil {
		return nil, err
	}

	if err = cursor.All(ctx, &plants); err != nil {
		return nil, err
	}

	return plants, nil
}

func GetItemInCollectionWithId(collectionName string, id string) (*models.Plant, error) {
	database := realClient.Database("test")
	collection := database.Collection(collectionName)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	objectId, convertErr := primitive.ObjectIDFromHex(id)
	var plant models.Plant

	if convertErr != nil {
		cancel()
		return nil, convertErr
	}

	err := collection.FindOne(ctx, bson.D{{"_id", objectId}}).Decode(&plant)
	cancel()
	if err != nil {
		return nil, err
	}

	return &plant, nil
}

func DeleteItemInCollection(collectionName string, id string) (int64, error) {
	database := realClient.Database("test")
	collection := database.Collection(collectionName)
	objectId, convertErr := primitive.ObjectIDFromHex(id)

	if convertErr != nil {
		return 0, convertErr
	}

	count, err := collection.DeleteOne(context.TODO(), bson.D{{"_id", objectId}})

	if err != nil {
		return count.DeletedCount, err
	}

	return count.DeletedCount, nil
}
