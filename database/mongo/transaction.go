package mongo

import (
	"time"

	"github.com/educolog9/helpers/enums"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Transaction struct {
	ID                    primitive.ObjectID   `bson:"_id"`
	OrganizationID        primitive.ObjectID   `bson:"organizationId"`
	PlanID                primitive.ObjectID   `bson:"planId"`
	Seats                 int                  `bson:"seats"`
	TotalPrice            float64              `bson:"totalPrice"`
	Categories            []primitive.ObjectID `bson:"categories"`
	OrderCode             string               `bson:"orderCode"`
	PaymentMethod         enums.PaymentMethod  `bson:"paymentMethod"`
	PaymentStatus         enums.PaymentStatus  `bson:"paymentStatus"`
	ExternalTransactionID string               `bson:"externalTransactionId"`
	Version               int                  `bson:"version"`
	CreatedAt             time.Time            `bson:"createdAt"`
	UpdatedAt             time.Time            `bson:"updatedAt"`
}
