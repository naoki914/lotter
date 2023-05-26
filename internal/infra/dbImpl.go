package infra

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/naoki914/lotter/internal/adapters"
	"github.com/naoki914/lotter/internal/domain"
	"github.com/naoki914/lotter/internal/domain/dhlotto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DbImpl struct {
	Client     *mongo.Client
	Database   *mongo.Database
	Collection *mongo.Collection
}

func createDrawInstance(lotteryType string) domain.Draw {
	switch lotteryType {
	case "dhlotto":
		return &dhlotto.DHDraw{}
	// add more cases as needed
	default:
		return nil
	}
}

func NewDbImpl(connectionString string, dbName string, collectionName string) adapters.DB {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(connectionString))
	if err != nil {
		fmt.Printf("Could not connect to mongo")
		os.Exit(2)
	}

	database := client.Database(dbName)
	collection := database.Collection(collectionName)

	return &DbImpl{
		Client:     client,
		Database:   database,
		Collection: collection,
	}
}

func (db *DbImpl) Create(draw domain.Draw) error {
	collection := db.Database.Collection("draws")
	_, err := collection.InsertOne(context.Background(), draw)
	return err
}

func (db *DbImpl) Get(id int) (*domain.Draw, error) {
	collection := db.Database.Collection("draws")
	var draw domain.Draw
	err := collection.FindOne(context.Background(), bson.M{"_id": id}).Decode(&draw)
	return &draw, err
}

func (db *DbImpl) GetAll(lotteryType string) ([]domain.Draw, error) {
	collection := db.Database.Collection("draws")
	cursor, err := collection.Find(context.TODO(), bson.D{{}})
	if err != nil {
		panic(err)
	}

	var bsonResults []bson.M
	if err = cursor.All(context.TODO(), &bsonResults); err != nil {
		fmt.Printf("could not parse results: %+v\n", err)
		os.Exit(201)
	}
	results := make([]domain.Draw, len(bsonResults))

	instance := createDrawInstance(lotteryType)
	if instance == nil {
		return nil, fmt.Errorf("unsupported lottery type: %s", lotteryType)
	}
	for i, bsonResult := range bsonResults {
		bsonBytes, _ := bson.Marshal(bsonResult)
		instance := createDrawInstance(lotteryType) // create new instance based on type
		bson.Unmarshal(bsonBytes, instance)
		results[i] = instance
	}

	return results, nil
}
