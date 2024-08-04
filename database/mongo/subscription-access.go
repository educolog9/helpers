package mongo

import (
	"time"

	"github.com/educolog9/helpers/enums"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type SubscriptionAccess struct {
	ID             primitive.ObjectID           `bson:"_id"`
	UserID         primitive.ObjectID           `bson:"userId"`
	OrganizationID primitive.ObjectID           `bson:"organizationId"`
	Categories     []primitive.ObjectID         `bson:"categories"`
	AccessType     enums.SubscriptionAccessType `bson:"accessType"`
	Version        int                          `bson:"version"`
	UpdatedAt      time.Time                    `bson:"updatedAt"`
}
