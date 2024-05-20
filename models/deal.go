package models

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Deal struct {
	Id        string    `json:"id"`
	Price     float64   `json:"price"`
	MaxItems  int       `json:"maxItems"`
	StartTime time.Time `json:"startTime"`
	EndTime   time.Time `json:"endTime"`
	ProductId int       `json:"productId"`
	Claimed   int       `json:"claimed"`
}



func (d *Deal) CreateDeal(db *mongo.Database) error {
	collection := db.Collection("deals")
	d.Id = primitive.NewObjectID().Hex()

	_, err := collection.InsertOne(context.TODO(), d)
	return err
}

func (d *Deal) ClaimDeal(db *mongo.Database, userID string) error {
	collection := db.Collection("deals")
	dealCollection := db.Collection("claimed-deals")

	filter := bson.M{"id": d.Id}

	claimFilter := bson.M{
		"deal_id": d.Id,
		"user_id": userID,
	}
	var session mongo.SessionContext
	claimCount, err := dealCollection.CountDocuments(session, claimFilter)
	if err != nil {
		log.Fatal(err)
	}

	if claimCount > 1 {
		return fmt.Errorf("User already claimed a deal")
	}

	if (d.Claimed > d.MaxItems)  {
		return fmt.Errorf("Deal is over")
	}
		

	update := bson.M{
		"$inc": bson.M{"claimed": 1},
	}

	_, err = collection.UpdateOne(session, filter, update)

	claim := bson.M{
		"deal_id": d.Id,
		"user_id": userID,
	}
	_, err = dealCollection.InsertOne(session, claim)
	if err != nil {
		session.AbortTransaction(session)
		return err
	}
	return session.CommitTransaction(session)
}

func (d *Deal) UpdateDeal(db *mongo.Database) error {
	collection := db.Collection("deals")
	filter := bson.M{"id": d.Id}
	update := bson.M{
		"$set": bson.M{
			"price":    d.Price,
			"maxItems": d.MaxItems,
			"endTime":  d.EndTime,
		},
	}
	_, err := collection.UpdateOne(context.TODO(), filter, update)
	return err
}

func (d *Deal) EndDeal(db *mongo.Database) error {
	collection := db.Collection("deals")
	filter := bson.M{"id": d.Id}
	update := bson.M{
		"$set": bson.M{
			"endTime": time.Now(),
		},
	}
	_, err := collection.UpdateOne(context.TODO(), filter, update)
	return err
}
