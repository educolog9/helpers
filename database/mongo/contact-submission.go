package mongo

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ContactSubmission struct {
	ID        primitive.ObjectID `bson:"_id"`
	Name      string             `bson:"name"`
	LastName  string             `bson:"lastName"`
	Email     string             `bson:"email"`
	Message   string             `bson:"message"`
	CreatedAt time.Time          `bson:"createdAt"`
	UpdatedAt time.Time          `bson:"updatedAt"`
}
