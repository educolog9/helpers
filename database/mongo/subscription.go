package mongo

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Subscription struct {
	ID             primitive.ObjectID `bson:"_id"`
	OrganizationID primitive.ObjectID `bson:"organizationId"`
	PaymentMethod  string             `bson:"paymentMethod"`
	ExternalID     string             `bson:"externalId"`
	PlanId         primitive.ObjectID `bson:"planId"`
	CreatedAt      time.Time          `bson:"createdAt"`
	NextPayment    bool               `bson:"nextPayment"`
	Version        int                `bson:"version"`
	UpdatedAt      time.Time          `bson:"updatedAt"`
}
