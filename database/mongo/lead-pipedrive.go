package mongo

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type LeadPipedrive struct {
	ID        primitive.ObjectID `bson:"_id"`
	Name      string             `bson:"name"`
	Email     string             `bson:"email"`
	Version   int                `bson:"version"`
	CreatedAt time.Time          `bson:"createdAt"`
	UpdatedAt time.Time          `bson:"updatedAt"`
}
