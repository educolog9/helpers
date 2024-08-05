package mongo

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Plan struct {
	ID            primitive.ObjectID `bson:"_id"`
	Price         float64            `bson:"price"`
	CategoryLimit int                `bson:"categoryLimit"`
	StripePlanID  string             `bson:"stripePlanID"`
	Status        string             `bson:"status"`
}
