package data

import (
	"context"
	"log"
	"time"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var client *mongo.Client

func New(mongo *mongo.Client) Models {
	client = mongo
	return Models{
		Location: Location{},
	}
}

type Models struct {
	Location Location	
}

type Location struct {
	ID		string   `bson:"_id,omitempty" json:"id,omitempty"`
	Name	string   `bson:"name" json:"name"`
	Description string `bson:"description" json:"description"`
}

func (l* Location) Insert(location Location) error {
	collection := client.Database("location").Collection("location")

	_, err := collection.InsertOne(context.TODO(), Location{
		Name: location.Name,
		Description: location.Description,
	})

	if err != nil {
		return err
	}
	return nil
}

func (l* Location) FindAll() ([]*Location, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	collection := client.Database("location").Collection("location")

	
	cursor, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		log.Println("Error finding all locations: ", err)
		return nil, err
	}

	defer cursor.Close(ctx)

	var locations []*Location
	
	for cursor.Next(ctx) {
		var location Location
		err = cursor.Decode(&location)
		if err != nil {
			log.Println("Error decoding location: ", err)
			return nil, err
		}
		locations = append(locations, &location)
	}

	return locations, nil
}

func (l* Location) FindById(id string) (*Location, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	collection := client.Database("location").Collection("location")

	docID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println("Error converting string to objectID: ", err)
		return nil, err
	}

	var location Location

	err = collection.FindOne(ctx, bson.M{"_id": docID}).Decode(&location)
	if err != nil {
		log.Println("Error finding location by ID: ", err)
		return nil, err
	}
	return &location, nil
}

func (l* Location) Update(location Location) error {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	collection := client.Database("location").Collection("location")

	docID, err := primitive.ObjectIDFromHex(location.ID)
	if err != nil {
		log.Println("Error converting string to objectID: ", err)
		return err
	}

	_, err = collection.UpdateOne(ctx, bson.M{"_id": docID}, bson.D{
		{"$set", bson.D{
			{"name", location.Name},
			{"description", location.Description},
		}},
	})
	
	if err != nil {
		log.Println("Error updating location: ", err)
		return err
	}

	return nil
}

func (l* Location) Delete(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	collection := client.Database("location").Collection("location")

	docID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println("Error converting string to objectID: ", err)
		return err
	}
	
	_, err = collection.DeleteOne(ctx, bson.M{"_id": docID})
	if err != nil {
		log.Println("Error deleting location: ", err)
		return err
	}

	return nil
}
	