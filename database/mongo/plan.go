package mongo

import (
	"time"

	"github.com/educolog9/helpers/enums"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Plan struct {
	ID                      primitive.ObjectID `bson:"_id"`
	StripePlanID            string             `bson:"stripePlanID"`
	Name                    string             `bson:"name"`
	NameTranslations        map[string]string  `bson:"Nametranslations"`
	Description             string             `bson:"description"`
	DescriptionTranslations map[string]string  `bson:"descriptionTranslations"`
	Price                   float64            `bson:"price"`
	CategoryLimit           int                `bson:"categoryLimit"`
	Version                 int                `bson:"version"`
	UpdatedAt               time.Time          `bson:"updatedAt"`
	Status                  enums.PlanStatus   `bson:"status"`
	CreatedAt               time.Time          `bson:"createdAt"`
}
