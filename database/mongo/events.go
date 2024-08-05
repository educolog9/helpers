package mongo

import (
	"time"

	"github.com/educolog9/helpers/enums"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Event struct {
	ID        primitive.ObjectID `bson:"_id"`
	EventID   string             `bson:"eventId"`
	Type      enums.EventType    `bson:"type"`
	Status    enums.LogStatus    `bson:"status"`
	CreatedAt time.Time          `bson:"createdAt"`
	UpdatedAt time.Time          `bson:"updatedAt"`
}
