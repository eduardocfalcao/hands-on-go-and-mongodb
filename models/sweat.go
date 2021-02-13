package models

import (
	"context"
	"fmt"
	"time"

	"github.com/eduardocfalcao/hands-on-go-and-mongodb/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const SWEAT_TABLE string = "sweat"

// Sweat - Sweat type
type Sweat struct {
	ID              primitive.ObjectID `bson:"_id,omitempty"`
	UserID          primitive.ObjectID `bson:"userid,omitempty" json:"userid"`
	CreatedAt       time.Time          `bson:"created_at"`
	Glucose         float32            `bson:"glucose" json:"glucose,omitempty"`
	Chloride        float32            `bson:"chloride" json:"chloride,omitempty"`
	Sodium          float32            `bson:"sodium" json:"sodium,omitempty"`
	Potassium       float32            `bson:"potassium" json:"potassium,omitempty"`
	Magnesium       float32            `bson:"magnesium" json:"magnesium,omitempty"`
	Calcium         float32            `bson:"calcium" json:"calcium,omitempty"`
	Humidity        float32            `bson:"humidity" json:"humidity,omitempty"`
	RoomTemperature float32            `bson:"room_temperature" json:"room_temperature,omitempty"`
	BodyTemperature float32            `bson:"body_temperature" json:"body_temperature,omitempty"`
	HeartBeat       int32              `bson:"heart_beat" json:"heart_beat,omitempty"`
}

// Create - Persist the Sweat object in the database
func (s *Sweat) Create() (err error) {
	db, err := db.GetDB()
	if err != nil {
		fmt.Println("No Database connection: ", err)
		return
	}

	s.CreatedAt = time.Now()
	collection := db.Collection(SWEAT_TABLE)
	_, err = collection.InsertOne(context.TODO(), s)
	if err != nil {
		fmt.Printf("Error Inserting sweat: %v", s)
		return
	}

	fmt.Println("Inserted sweat into collection")
	return
}

// ListAllSweat - lists all sweats from database
func ListAllSweat() (sweats []Sweat, err error) {
	db, err := db.GetDB()
	if err != nil {
		fmt.Println("No Database connection: ", err)
		return
	}

	collection := db.Collection(SWEAT_TABLE)
	ctx := context.TODO()
	cursor, err := collection.Find(ctx, bson.D{})
	if err != nil {
		return
	}
	defer cursor.Close(ctx)

	var elem Sweat
	for cursor.Next(ctx) {
		err = cursor.Decode(&elem)
		sweats = append(sweats, elem)
	}
	if err = cursor.Err(); err != nil {
		fmt.Println("Error in listing data: ", err)
		return
	}
	return
}

// GetSweatByID returns a Sweat by it's ID
func GetSweatByID(id string) (sweat Sweat, err error) {
	db, err := db.GetDB()
	if err != nil {
		fmt.Println("No database connection: ", err)
		return
	}

	collection := db.Collection(SWEAT_TABLE)
	ctx := context.TODO()
	objectID, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		fmt.Println("Error converting string to ObjectID. Error: ", err)
		return
	}

	result := collection.FindOne(ctx, bson.M{"_id": objectID})

	if err = result.Decode(&sweat); err != nil {
		fmt.Println("Error decoding Sweat from database: ", err)
		return
	}
	return
}
