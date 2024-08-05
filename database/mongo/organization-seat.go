package mongo

import (
	"time"

	"github.com/educolog9/helpers/enums"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type OrganizationSeat struct {
	ID             primitive.ObjectID           `bson:"_id"`
	OrganizationID primitive.ObjectID           `bson:"organizationId"`
	Seats          int                          `bson:"seats"`
	AccessType     enums.SubscriptionAccessType `bson:"accessType"`
	CreatedAt      time.Time
	SeatsPurchased int       `bson:"seatsPurchased"`
	Version        int       `bson:"version"`
	UpdatedAt      time.Time `bson:"updatedAt"`
}
