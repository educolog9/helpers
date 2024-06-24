package mongo

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type HoursManagement struct {
	ID             primitive.ObjectID `bson:"_id"`
	OrganizationID primitive.ObjectID `bson:"organizationId"`
	CategoryID     primitive.ObjectID `bson:"categoryId"`
	Year           int                `bson:"year"`
	Month          int                `bson:"month"`
	EngagementTime int                `bson:"engagementTime"`
	CreatedAt      time.Time          `bson:"createdAt"`
	UpdatedAt      time.Time          `bson:"updatedAt"`
}
